package router

import (
	"net/http"

	controller "github.com/DevIdol/Golang_JWT/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controller.AuthController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome")
	})

	router := service.Group("/api")

	authRouter := router.Group("/auth")
	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)

	return service
}
