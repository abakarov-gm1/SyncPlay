package app

import (
	"goWin/services/auth-service/internal/repository"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	UserService     *UserService
	RegisterService *RegisterService
}

func NewApp(db *gorm.DB) *ServiceContainer {
	userRepo := repository.NewUserRepository(db)
	userService := NewUserService(userRepo)
	registerService := NewRegisterService(userRepo)

	return &ServiceContainer{
		UserService:     userService,
		RegisterService: registerService,
	}

}
