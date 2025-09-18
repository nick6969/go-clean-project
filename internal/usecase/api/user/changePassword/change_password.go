package changePassword

import (
	"context"
	"errors"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type repository interface {
	FindUserByID(ctx context.Context, userID int) (*domain.DBUserModel, error)
	UpdateUserPassword(ctx context.Context, user *domain.DBUserModel) error
}

type password interface {
	Compare(hashed, password string) error
	Hash(password string) (string, error)
}

type UseCase struct {
	repository repository
	password   password
}

func NewUseCase(repository repository, password password) *UseCase {
	return &UseCase{repository: repository, password: password}
}

type Input struct {
	userID      int
	oldPassword string
	newPassword string
}

func NewInput(userID int, oldPassword string, newPassword string) Input {
	return Input{
		userID:      userID,
		oldPassword: oldPassword,
		newPassword: newPassword,
	}
}

func (u *UseCase) Execute(ctx context.Context, input Input) error {
	user, err := u.repository.FindUserByID(ctx, input.userID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	if err := u.password.Compare(user.PasswordHash(), input.oldPassword); err != nil {
		return errors.New("invalid old password")
	}

	hashedPassword, err := u.password.Hash(input.newPassword)
	if err != nil {
		return err
	}

	if err := user.ChangePassword(hashedPassword); err != nil {
		return err
	}

	if err := u.repository.UpdateUserPassword(ctx, user); err != nil {
		return err
	}

	return nil
}
