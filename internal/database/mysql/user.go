package mysql

import (
	"context"
	"errors"

	"github.com/nick6969/go-clean-project/internal/database/mysql/entity"
	"github.com/nick6969/go-clean-project/internal/domain"
	"gorm.io/gorm"
)

func (d *Database) CheckEmailIsExists(ctx context.Context, email string) (bool, *domain.GPError) {
	var count int64
	err := d.db.WithContext(ctx).
		Model(entity.User{}).
		Where("email = ?", email).
		Count(&count).Error
	if err != nil {
		return false, domain.NewGPErrorWithError(domain.ErrCodeDatabaseError, err).Append("checking email exists")
	}

	return count > 0, nil
}

func (d *Database) CreateUser(ctx context.Context, email, hashedPassword string) (int, *domain.GPError) {
	user := &entity.User{
		Email:    email,
		Password: hashedPassword,
	}

	if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
		return 0, domain.NewGPErrorWithError(domain.ErrCodeDatabaseError, err).Append("creating user")
	}

	return user.ID, nil
}

func (d *Database) FindUserByEmail(ctx context.Context, email string) (*domain.DBUserModel, *domain.GPError) {
	var user entity.User
	err := d.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, domain.NewGPErrorWithError(domain.ErrCodeDatabaseError, err).Append("finding user by email")
	}

	return user.ToDomain()
}

func (d *Database) FindUserByID(ctx context.Context, userID int) (*domain.DBUserModel, *domain.GPError) {
	var user entity.User
	err := d.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", userID).
		First(&user).Error
	if err != nil {
		return nil, domain.NewGPErrorWithError(domain.ErrCodeDatabaseError, err).Append("finding user by id")
	}

	return user.ToDomain()
}

func (d *Database) UpdateUserPassword(ctx context.Context, user *domain.DBUserModel) *domain.GPError {
	err := d.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Update("password", user.PasswordHash()).Error
	if err != nil {
		return domain.NewGPErrorWithError(domain.ErrCodeDatabaseError, err).Append("updating user password")
	}
	return nil
}
