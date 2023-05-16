package controller

import (
	"net/http"

	"github.com/DevIdol/Golang_JWT/data/response"
	"github.com/DevIdol/Golang_JWT/respository"
	"github.com/gin-gonic/gin"
)

type Usercontroller struct {
	UserRespository respository.UserRespository
}

func NewUsercontroller(respository respository.UserRespository) *Usercontroller {
	return &Usercontroller{UserRespository: respository}
}

func (controller *Usercontroller) GetUsers(ctx *gin.Context) {
	users := controller.UserRespository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Get All Users Success",
		Data:    users,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
