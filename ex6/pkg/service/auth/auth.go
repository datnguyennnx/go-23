package auth

import (
	"errors"

	"github.com/datnguyennnx/go-23/ex6/pkg/components/tokenprovider"
	"github.com/datnguyennnx/go-23/ex6/pkg/components/tokenprovider/jwt"
	"github.com/datnguyennnx/go-23/ex6/pkg/config"
	request "github.com/datnguyennnx/go-23/ex6/pkg/data/request/user"
	"github.com/datnguyennnx/go-23/ex6/pkg/model"
	"github.com/datnguyennnx/go-23/ex6/pkg/repo"
	"github.com/datnguyennnx/go-23/ex6/pkg/service"
	"github.com/datnguyennnx/go-23/ex6/pkg/utils"
	"github.com/go-playground/validator"
)

type AuthenticationServiceImpl struct {
	UsersRepository repo.UserRepo
	Validate        *validator.Validate
}

func NewAuthenticationService(usersRepository repo.UserRepo, validate *validator.Validate) service.AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (*tokenprovider.Token, error) {
	new_users, users_err := a.UsersRepository.User().FindByEmail(users.Email)
	if users_err != nil {
		return nil, errors.New("invalid username or Password 1")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return nil, errors.New("invalid username or Password 2")
	}

	// Generate Token
	tokenProvider := jwt.NewTokenJWTProvider(config.AccessTokenPrivateKey)

	payload := tokenprovider.TokenPayload{
		ID:   new_users.ID,
		Role: new_users.Role,
	}

	accessToken, err := tokenProvider.Generate(payload, int(config.AccessTokenExpiresIn))
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (a *AuthenticationServiceImpl) Register(users request.SignUpInput) {
	hashedPassword, err := utils.HashPassword(users.Password)
	if err != nil {
		panic(err)
	}
	newUser := model.Users{
		Name:     users.Name,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UsersRepository.User().Save(newUser)
}
