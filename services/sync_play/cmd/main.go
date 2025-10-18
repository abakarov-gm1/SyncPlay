package main

import (
	"github.com/joho/godotenv"
	"goWin/services/syncPlay/internal/app"
	"goWin/services/syncPlay/internal/database"
	"goWin/services/syncPlay/internal/route"
	"goWin/services/syncPlay/models"
	"log"
)

func main() {
	_ = godotenv.Load()

	database.Init()

	if err := database.DB.AutoMigrate(&models.Room{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	services := app.NewApp(database.DB)
	route.Server(services)
}
