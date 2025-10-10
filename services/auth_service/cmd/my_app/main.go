package main

import (
	"github.com/joho/godotenv"
	"goWin/services/auth-service/internal/app"
	"goWin/services/auth-service/internal/database"
	"goWin/services/auth-service/internal/models"
	"goWin/services/auth-service/internal/route"
	"log"
)

func main() {
	_ = godotenv.Load()

	database.Init()

	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	serviceContainer := app.NewApp(database.DB)

	err := route.Server(serviceContainer)
	if err != nil {
		panic(err)
	}

}
