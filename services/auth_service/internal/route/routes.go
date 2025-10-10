package route

import (
	"github.com/go-chi/chi/v5"
	"goWin/services/auth-service/internal/app"
	"goWin/services/auth-service/internal/route/handlers"
)

func RegisterRoutes(r *chi.Mux, serviceContainer *app.ServiceContainer) {
	h := handlers.NewHandlerContainer(serviceContainer)

	r.Post("/login", h.AuthHandler.Login)
	r.Post("/register", h.AuthHandler.Register)
	r.Get("/users", h.UserHandler.GetAll)
}
