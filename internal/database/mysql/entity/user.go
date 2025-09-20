package entity

import (
	"time"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type User struct {
	ID        int       `gorm:"column:id"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password_hash"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) ToDomain() (*domain.DBUserModel, *domain.GPError) {
	user, err := domain.NewDBUserModel(u.ID, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return nil, domain.NewGPErrorWithError(domain.ErrCodeModelConvertError, err).Append("failed to convert User entity to DBUserModel")
	}
	return user, nil
}

func NewUserFromDomain(m *domain.DBUserModel) *User {
	return &User{
		ID:        m.ID(),
		Email:     m.Email(),
		Password:  m.PasswordHash(),
		CreatedAt: m.CreatedAt(),
		UpdatedAt: m.UpdatedAt(),
	}
}
