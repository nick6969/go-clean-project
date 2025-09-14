package login

import (
	"context"
	"fmt"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type repository interface {
	FindUserByEmail(ctx context.Context, email string) (*domain.DBUserModel, error)
}

type password interface {
	Compare(hashed, password string) error
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
	// 1. 根據 email 從資料庫取得 user 資料
	user, err := u.repository.FindUserByEmail(ctx, input.email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	// 2. 比對找出使用者的密碼跟輸入的密碼是否相同
	if err := u.password.Compare(user.PasswordHash(), input.password); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	// 3. 產生 access Token
	accessToken, err := u.token.GenerateAccessToken(user.ID())
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// 4. 回傳 access Token
	return &Output{
		AccessToken: accessToken,
	}, nil
}
