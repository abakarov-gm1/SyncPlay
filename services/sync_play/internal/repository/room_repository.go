package repository

import (
	"goWin/services/syncPlay/models"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) CreateRoom(repo *models.Room) error {
	return r.db.Create(repo).Error
}

func (r *RoomRepository) GetRooms() ([]models.Room, error) {
	var rooms []models.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}
