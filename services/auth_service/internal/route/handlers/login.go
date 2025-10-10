package handlers

import (
	"encoding/json"
	"goWin/services/auth-service/internal/app"
	"goWin/services/auth-service/internal/drt"
	"goWin/services/auth-service/internal/models"
	"net/http"
	"time"
)

type AuthHandler struct {
	services *app.ServiceContainer
}

func NewAuthHandler(services *app.ServiceContainer) *AuthHandler {
	return &AuthHandler{services: services}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData drt.Login
	var loginResponse drt.ResponseLogin
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		loginResponse.Text = "Данные не валидны"
		loginResponse.Status = false
		_ = json.NewEncoder(w).Encode(loginResponse)
	}
	response, _ := h.services.RegisterService.Login(&loginData)

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    response.Token,
		HttpOnly: true,
		Secure:   false, // true в продакшене
		SameSite: http.SameSiteLaxMode,
		Domain:   ".localhost",
		Path:     "/",
		MaxAge:   15 * 24 * 3600,
		Expires:  time.Now().Add(15 * 24 * time.Hour),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)

}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerData models.User
	if err := json.NewDecoder(r.Body).Decode(&registerData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	if err := h.services.RegisterService.Register(&registerData); err != nil {
		_, _ = w.Write([]byte(err.Error()))
	} else {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Пользователь успешно зарегистрирован",
		})
	}
}
