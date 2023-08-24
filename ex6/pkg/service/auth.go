package service

import (
	"github.com/datnguyennnx/go-23/ex6/pkg/components/tokenprovider"
	request "github.com/datnguyennnx/go-23/ex6/pkg/data/request/user"
)

type AuthenticationService interface {
	Login(user request.LoginRequest) (*tokenprovider.Token, error)
	Register(user request.SignUpInput)
}
