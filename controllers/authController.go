package controller

import (
	"net/http"

	"github.com/DevIdol/Golang_JWT/data/request"
	"github.com/DevIdol/Golang_JWT/data/response"
	"github.com/DevIdol/Golang_JWT/helpers"
	service "github.com/DevIdol/Golang_JWT/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.Authservice
}

func NewAuthController(service service.Authservice) *AuthController {
	return &AuthController{
		AuthService: service,
	}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helpers.ErrorPanic(err)

	token, err_token := controller.AuthService.Login(loginRequest)

	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid Email or Password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Login Success",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helpers.ErrorPanic(err)

	controller.AuthService.Register(createUserRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Created User Success",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
