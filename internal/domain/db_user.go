package domain

import (
	"errors"
	"time"
)

type DBUserModel struct {
	id           int
	email        string
	passwordHash string
	createdAt    time.Time
	updatedAt    time.Time
}

func NewDBUserModel(id int, email, passwordHash string, createdAt, updatedAt time.Time) (*DBUserModel, error) {
	if id < 0 {
		return nil, errors.New("id must be non-negative")
	}

	if email == "" {
		return nil, errors.New("email is required")
	}

	if passwordHash == "" {
		return nil, errors.New("passwordHash is required")
	}

	return &DBUserModel{
		id:           id,
		email:        email,
		passwordHash: passwordHash,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}, nil
}

func (u *DBUserModel) ID() int {
	return u.id
}

func (u *DBUserModel) Email() string {
	return u.email
}

func (u *DBUserModel) PasswordHash() string {
	return u.passwordHash
}

func (u *DBUserModel) CreatedAt() time.Time {
	return u.createdAt
}

func (u *DBUserModel) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *DBUserModel) ChangePassword(newPasswordHash string) error {
	if newPasswordHash == "" {
		return errors.New("new password hash is required")
	}
	u.passwordHash = newPasswordHash
	return nil
}
