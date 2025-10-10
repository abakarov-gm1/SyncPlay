package app

import (
	"goWin/services/auth-service/internal/models"
	"goWin/services/auth-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetAll() ([]models.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
