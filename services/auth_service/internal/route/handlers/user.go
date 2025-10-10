package handlers

import (
	"encoding/json"
	"goWin/services/auth-service/internal/app"
	"net/http"
)

type UserHandler struct {
	services *app.ServiceContainer
}

func NewUserHandler(services *app.ServiceContainer) *UserHandler {
	return &UserHandler{services: services}
}

func (s *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := s.services.UserService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
