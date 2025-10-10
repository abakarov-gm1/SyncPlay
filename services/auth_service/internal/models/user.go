package models

import "gorm.io/gorm"

type User struct {
	gorm.Model         // включает ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string  `gorm:"size:100;not null"`
	photo      *string `gorm:"size:255"`
	Password   string  `gorm:"size:100;not null"`
	Status     string  `gorm:"not null"`
}
