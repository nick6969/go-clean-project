package changePassword

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type repository interface {
	FindUserByID(ctx context.Context, userID int) (*domain.DBUserModel, *domain.GPError)
	UpdateUserPassword(ctx context.Context, user *domain.DBUserModel) *domain.GPError
}

type password interface {
	Compare(hashed, password string) *domain.GPError
	Hash(password string) (string, *domain.GPError)
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

func (u *UseCase) Execute(ctx context.Context, input Input) *domain.GPError {
	user, err := u.repository.FindUserByID(ctx, input.userID)
	if err != nil {
		return err.Append("failed to find user by id")
	}

	if user == nil {
		return domain.NewGPError(domain.ErrCodeUserNotFound)
	}

	if err := u.password.Compare(user.PasswordHash(), input.oldPassword); err != nil {
		return err.Append("change password failed")
	}

	hashedPassword, err := u.password.Hash(input.newPassword)
	if err != nil {
		return err.Append("change password failed")
	}

	if err := user.ChangePassword(hashedPassword); err != nil {
		return domain.NewGPError(domain.ErrCodeParametersNotCorrect).Append("change password failed")
	}

	if err := u.repository.UpdateUserPassword(ctx, user); err != nil {
		return err.Append("change password failed")
	}

	return nil
}
