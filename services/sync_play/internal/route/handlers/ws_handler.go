package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"goWin/services/syncPlay/internal/app"
	"goWin/services/syncPlay/internal/drt"
	"net/http"
)

type WsHandler struct {
	services *app.ServiceContainer
}

func NewWsHandler(services *app.ServiceContainer) *WsHandler {
	return &WsHandler{services: services}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func (ws *WsHandler) Ws(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WsHandler.Ws triggered")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true
	defer delete(clients, conn)

	for {
		var msg drt.Message
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
