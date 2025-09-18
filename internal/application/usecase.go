package application

import (
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/changePassword"
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/login"
	"github.com/nick6969/go-clean-project/internal/usecase/api/user/register"
)

type UseCase struct {
	User *UserUseCase
}

func NewUseCase(app *Application) *UseCase {
	return &UseCase{
		User: NewUserUseCase(app),
	}
}

type UserUseCase struct {
	Register       *register.UseCase
	Login          *login.UseCase
	ChangePassword *changePassword.UseCase
}

func NewUserUseCase(app *Application) *UserUseCase {
	return &UserUseCase{
		Register:       register.NewUseCase(app.Database, app.Service.Password, app.Service.Token),
		Login:          login.NewUseCase(app.Database, app.Service.Password, app.Service.Token),
		ChangePassword: changePassword.NewUseCase(app.Database, app.Service.Password),
	}
}
