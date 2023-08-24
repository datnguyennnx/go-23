package main

import (
	"log"
	"net/http"
	"time"

	"github.com/datnguyennnx/go-23/ex6/pkg/config"
	"github.com/datnguyennnx/go-23/ex6/pkg/controller"
	"github.com/datnguyennnx/go-23/ex6/pkg/model"
	"github.com/datnguyennnx/go-23/ex6/pkg/repo"
	"github.com/datnguyennnx/go-23/ex6/pkg/router"
	service "github.com/datnguyennnx/go-23/ex6/pkg/service/auth"
	"github.com/go-playground/validator"
)

func main() {
	init, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	validate := validator.New()
	//Init Repository
	userRepository := repo.NewPGRepo(config.ConnectStringDB(&init))
	userRepository.AutoMigrate(&model.Users{})

	//Init Service
	authenticationService := service.NewAuthenticationService(userRepository, validate)
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)

	routes := router.NewRouter(userRepository, authenticationController, usersController)

	server := &http.Server{
		Addr:           ":" + init.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server_err := server.ListenAndServe()
	panic(server_err)
}
