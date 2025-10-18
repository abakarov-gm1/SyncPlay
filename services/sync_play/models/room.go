package models

type Room struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Link       string `gorm:"type:text"`
	UserIDRoot uint   `gorm:"not null"`  // создатель комнаты
	UserIDs    []uint `gorm:"-"`         // просто временное поле, не хранится в БД
	UsersJSON  string `gorm:"type:json"` // JSON массив ID пользователей
}
