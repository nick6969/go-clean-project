package application

import "github.com/nick6969/go-clean-project/internal/service/password"

type Service struct {
	Password *password.Service
}

func NewService(app *Application) *Service {
	return &Service{
		Password: password.NewService(),
	}
}
