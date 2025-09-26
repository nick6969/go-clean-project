package register

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type repository interface {
	CheckEmailIsExists(ctx context.Context, email string) (bool, *domain.GPError)
	CreateUser(ctx context.Context, email, hashedPassword string) (int, *domain.GPError)
}

type password interface {
	Hash(password string) (string, *domain.GPError)
}

type token interface {
	GenerateAccessToken(userID int) (string, *domain.GPError)
}

type dispatcher interface {
	DispatchUserRegistered(ctx context.Context, userID int, email string)
}

type UseCase struct {
	repository repository
	password   password
	token      token
	dispatcher dispatcher
}

func NewUseCase(repository repository, password password, token token, dispatcher dispatcher) *UseCase {
	return &UseCase{repository: repository, password: password, token: token, dispatcher: dispatcher}
}

type Input struct {
	email    string
	password string
}

func NewInput(email, password string) Input {
	return Input{
		email:    email,
		password: password,
	}
}

type Output struct {
	AccessToken string
}

func (u *UseCase) Execute(ctx context.Context, input Input) (*Output, *domain.GPError) {
	isExists, err := u.repository.CheckEmailIsExists(ctx, input.email)
	if err != nil {
		return nil, err.Append("register")
	}

	if isExists {
		return nil, domain.NewGPError(domain.ErrCodeUserEmailExists)
	}

	hashedPassword, err := u.password.Hash(input.password)
	if err != nil {
		return nil, err.Append("register")
	}

	userID, err := u.repository.CreateUser(ctx, input.email, hashedPassword)
	if err != nil {
		return nil, err.Append("register")
	}

	u.dispatcher.DispatchUserRegistered(ctx, userID, input.email)

	accessToken, err := u.token.GenerateAccessToken(userID)
	if err != nil {
		return nil, err.Append("register")
	}

	return &Output{
		AccessToken: accessToken,
	}, nil
}
