package app

import (
	"goWin/services/auth-service/internal/auth"
	"goWin/services/auth-service/internal/drt"
	"goWin/services/auth-service/internal/models"
	"goWin/services/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	repo *repository.UserRepository
}

func NewRegisterService(repo *repository.UserRepository) *RegisterService {
	return &RegisterService{repo}
}

func (r *RegisterService) Register(newUser *models.User) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)
	newUser.Password = string(hashPassword)
	return r.repo.Create(newUser)
}

func (r *RegisterService) Login(loginData *drt.Login) (*drt.ResponseLogin, error) {
	user, err := r.repo.GetUser(loginData.Name)
	resp := &drt.ResponseLogin{Text: ""}
	if err != nil || user == nil {
		resp.Text = "Пользователя не существует !!!"
		return resp, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		resp.Text = "Пароли не совпадают !!!"
		resp.Status = false
		return resp, err
	}

	tokenString, err := auth.GenerateToken(int(user.ID))

	if err != nil {
		resp.Text = err.Error()
		resp.Status = false
	}

	resp.Text = "ok"
	resp.Status = true
	resp.Token = tokenString
	return resp, nil
}
