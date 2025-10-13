package handlers

import "goWin/services/syncPlay/internal/app"

type HandlerContainer struct {
	WsHandler *WsHandler
}

func NewHandlerRegister(services *app.ServiceContainer) *HandlerContainer {

	return &HandlerContainer{
		WsHandler: NewWsHandler(services),
	}
}
