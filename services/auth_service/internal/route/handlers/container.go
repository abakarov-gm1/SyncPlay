package handlers

import "goWin/services/auth-service/internal/app"

type HandlerContainer struct {
	AuthHandler *AuthHandler
	UserHandler *UserHandler
}

func NewHandlerContainer(services *app.ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		AuthHandler: NewAuthHandler(services),
		UserHandler: NewUserHandler(services),
	}
}
