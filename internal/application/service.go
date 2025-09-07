package application

import (
	"github.com/nick6969/go-clean-project/internal/service/password"
	"github.com/nick6969/go-clean-project/internal/service/token"
)

type Service struct {
	Password *password.Service
	Token    *token.Service
}

func NewService(app *Application) (*Service, error) {
	passwordService := password.NewService()
	tokenService, err := token.NewService([]byte(app.Config.Token.Secret))
	if err != nil {
		return nil, err
	}

	return &Service{
		Password: passwordService,
		Token:    tokenService,
	}, nil
}
