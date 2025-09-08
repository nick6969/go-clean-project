package register

import (
	"context"
	"errors"
)

type repository interface {
	CheckEmailIsExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, email, hashedPassword string) (int, error)
}

type password interface {
	Hash(password string) (string, error)
}

type token interface {
	GenerateAccessToken(userID int) (string, error)
}

type UseCase struct {
	repository repository
	password   password
	token      token
}

func NewUseCase(repository repository, password password, token token) *UseCase {
	return &UseCase{repository: repository, password: password, token: token}
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

func (u *UseCase) Execute(ctx context.Context, input Input) (*Output, error) {
	isExists, err := u.repository.CheckEmailIsExists(ctx, input.email)
	if err != nil {
		return nil, err
	}

	if isExists {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := u.password.Hash(input.password)
	if err != nil {
		return nil, err
	}

	userID, err := u.repository.CreateUser(ctx, input.email, hashedPassword)
	if err != nil {
		return nil, err
	}

	accessToken, err := u.token.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	return &Output{
		AccessToken: accessToken,
	}, nil
}
