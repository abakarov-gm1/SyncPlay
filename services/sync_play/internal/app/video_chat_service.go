package app

import (
	"goWin/services/syncPlay/internal/repository"
)

type VideoChatService struct {
	repo *repository.RoomRepository
}

func NewVideoChatService(repo *repository.RoomRepository) *VideoChatService {
	return &VideoChatService{repo: repo}
}
