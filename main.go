package main

import (
	"log"
	"net/http"

	"github.com/DevIdol/Golang_JWT/config"
	controller "github.com/DevIdol/Golang_JWT/controllers"
	"github.com/DevIdol/Golang_JWT/helpers"
	"github.com/DevIdol/Golang_JWT/models"
	"github.com/DevIdol/Golang_JWT/respository"
	"github.com/DevIdol/Golang_JWT/router"
	service "github.com/DevIdol/Golang_JWT/services"
	"github.com/go-playground/validator/v10"
)

func main() {
	loadENV, err := config.LoadENV(".")

	if err != nil {
		log.Fatal("Could not load env variables", err)
	}

	//Database
	db, _ := config.ConnectDB(&loadENV)
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.User{})

	userRespository := respository.NewUserRespositoryImpl(db)

	authService := service.NewAuthServiceImpl(userRespository, validate)

	authController := controller.NewAuthController(authService)

	routes := router.NewRouter(authController)

	server := http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helpers.ErrorPanic(server_err)
}
