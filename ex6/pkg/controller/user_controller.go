package controller

import (
	"net/http"

	response "github.com/datnguyennnx/go-23/ex6/pkg/data/response/user"
	"github.com/datnguyennnx/go-23/ex6/pkg/repo"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repo.UserRepo
}

func NewUsersController(repository repo.UserRepo) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	users, err := controller.userRepository.User().FindAll()
	if err != nil {
		panic(err)
	}
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
