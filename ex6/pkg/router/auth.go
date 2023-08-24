package router

import (
	"net/http"

	"github.com/datnguyennnx/go-23/ex6/pkg/controller"
	"github.com/datnguyennnx/go-23/ex6/pkg/repo"
	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repo.UserRepo, authenticationController *controller.AuthenticationController, usersController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	return service
}
