package route

import (
	"github.com/go-chi/chi/v5"
	"goWin/services/auth-service/internal/app"
	"net/http"
)

func Server(serviceContainer *app.ServiceContainer) error {
	r := chi.NewRouter()

	RegisterRoutes(r, serviceContainer)

	return http.ListenAndServe(":8000", r)
}
