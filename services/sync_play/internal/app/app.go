package app

import (
	"goWin/services/syncPlay/internal/repository"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	VideoChat *VideoChatService
}

func NewApp(db *gorm.DB) *ServiceContainer {
	RoomRepo := repository.NewRoomRepository(db)
	wsService := NewVideoChatService(RoomRepo)

	return &ServiceContainer{
		VideoChat: wsService,
	}
}
