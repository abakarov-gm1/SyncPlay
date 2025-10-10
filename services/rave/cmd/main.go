package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Message struct {
	Type    string  `json:"type"` // "chat" или "action"
	Message string  `json:"message,omitempty"`
	State   string  `json:"state,omitempty"` // "paused"/"playing"
	Time    float64 `json:"time,omitempty"`  // позиция видео
	LocalId string  `json:"localId"`         // уникальный ID клиента
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true
	defer delete(clients, conn)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading JSON:", err)
			break
		}

		fmt.Println("Received:", msg)

		// Рассылаем всем, кроме отправителя
		for client := range clients {
			if client != conn && msg.Message == "" {
				if err := client.WriteJSON(msg); err != nil {
					client.Close()
					delete(clients, client)
				}
			} else {
				if err := client.WriteJSON(msg); err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
