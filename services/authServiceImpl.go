package service

import (
	"errors"

	"github.com/DevIdol/Golang_JWT/config"
	"github.com/DevIdol/Golang_JWT/data/request"
	"github.com/DevIdol/Golang_JWT/helpers"
	"github.com/DevIdol/Golang_JWT/models"
	"github.com/DevIdol/Golang_JWT/respository"
	"github.com/DevIdol/Golang_JWT/utils"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UserRespository respository.UserRespository
	Validate        *validator.Validate
}

func NewAuthServiceImpl(usersRespository respository.UserRespository, validate *validator.Validate) Authservice {
	return &AuthServiceImpl{
		UserRespository: usersRespository,
		Validate:        validate,
	}
}

// Login implements Authservice
func (auth *AuthServiceImpl) Login(users request.LoginRequest) (string, error) {

	new_user, err := auth.UserRespository.FindByUsername(users.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	config, _ := config.LoadENV(".")
	verufy_password := utils.VerifyPassword(new_user.Password, users.Password)

	if verufy_password != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(config.TokenExpireIn, new_user.Id, config.TokenSecret)
	helpers.ErrorPanic(err)
	return token, nil

}

// Register implements Authservice
func (auth *AuthServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helpers.ErrorPanic(err)

	newUser := models.User{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	auth.UserRespository.Save(newUser)

}
