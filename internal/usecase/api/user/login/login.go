package login

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type repository interface {
	FindUserByEmail(ctx context.Context, email string) (*domain.DBUserModel, *domain.GPError)
}

type password interface {
	Compare(hashed, password string) *domain.GPError
}

type token interface {
	GenerateAccessToken(userID int) (string, *domain.GPError)
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

func (u *UseCase) Execute(ctx context.Context, input Input) (*Output, *domain.GPError) {
	// 1. 根據 email 從資料庫取得 user 資料
	user, err := u.repository.FindUserByEmail(ctx, input.email)
	if err != nil {
		return nil, err.Append("login")
	}

	if user == nil {
		return nil, domain.NewGPError(domain.ErrCodeUserNotFound)
	}

	// 2. 比對找出使用者的密碼跟輸入的密碼是否相同
	if err := u.password.Compare(user.PasswordHash(), input.password); err != nil {
		return nil, err.Append("login")
	}

	// 3. 產生 access Token
	accessToken, err := u.token.GenerateAccessToken(user.ID())
	if err != nil {
		return nil, err.Append("login")
	}

	// 4. 回傳 access Token
	return &Output{
		AccessToken: accessToken,
	}, nil
}
