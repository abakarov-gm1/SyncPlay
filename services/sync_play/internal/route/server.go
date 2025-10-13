package route

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"goWin/services/syncPlay/internal/app"
	"net/http"
)

func Server(services *app.ServiceContainer) {

	ch := chi.NewRouter()

	RegisterHandler(ch, services)

	//h := handlers.NewHandlerRegister(services)
	//http.HandleFunc("/ws", h.WsHandler.Ws)

	if err := http.ListenAndServe(":8080", ch); err != nil {
		fmt.Println("Server error:", err)
	}
}
