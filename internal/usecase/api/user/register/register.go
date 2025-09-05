package register

import (
	"context"
)

type repository interface {
	CheckEmailIsExists(email string) (bool, error)
	CreateUser(email, hashedPassword string) (int, error)
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
	// 業務邏輯處理
	// 與 repository 或 service 互動
	return &Output{}, nil
}
