package main

import (
	"log"
	"net/http"

	"github.com/datnguyennnx/go-23/ex6/pkg/config"
	"github.com/datnguyennnx/go-23/ex6/pkg/model"
	"github.com/datnguyennnx/go-23/ex6/pkg/repo"
	"github.com/gin-gonic/gin"
)

func initServer(cfg config.Config) *gin.Engine {
	srv := gin.Default()

	srv.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	return srv
}

func main() {
	init, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	r := repo.NewPGRepo(config.ConnectStringDB(&init))
	r.AutoMigrate(&model.Users{})

	srv := initServer(init)
	srv.Run(":" + init.ServerPort)
}
