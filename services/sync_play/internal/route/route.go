package route

import (
	"github.com/go-chi/chi/v5"
	"goWin/services/syncPlay/internal/app"
	"goWin/services/syncPlay/internal/route/handlers"
)

func RegisterHandler(r *chi.Mux, services *app.ServiceContainer) {
	h := handlers.NewHandlerRegister(services)
	r.HandleFunc("/ws", h.WsHandler.Ws)
}
