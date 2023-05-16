package service

import "github.com/DevIdol/Golang_JWT/data/request"

type Authservice interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
}
