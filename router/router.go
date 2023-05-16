package router

import (
	"net/http"

	controller "github.com/DevIdol/Golang_JWT/controllers"
	"github.com/DevIdol/Golang_JWT/middleware"
	"github.com/DevIdol/Golang_JWT/respository"
	"github.com/gin-gonic/gin"
)

func NewRouter(userRespository respository.UserRespository, authController *controller.AuthController, userController *controller.Usercontroller) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome")
	})

	router := service.Group("/api")

	authRouter := router.Group("/auth")
	authRouter.POST("/register", authController.Register)
	authRouter.POST("/login", authController.Login)

	userRouter := router.Group("/users")
	userRouter.GET("", middleware.IsAuth(userRespository), userController.GetUsers)

	return service
}
