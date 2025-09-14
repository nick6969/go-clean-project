package mysql

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/database/mysql/entity"
	"github.com/nick6969/go-clean-project/internal/domain"
)

func (d *Database) CheckEmailIsExists(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := d.db.WithContext(ctx).
		Model(entity.User{}).
		Where("email = ?", email).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (d *Database) CreateUser(ctx context.Context, email, hashedPassword string) (int, error) {
	user := &entity.User{
		Email:    email,
		Password: hashedPassword,
	}

	if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (d *Database) FindUserByEmail(ctx context.Context, email string) (*domain.DBUserModel, error) {
	var user entity.User
	if err := d.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	return domain.NewDBUserModel(user.ID, user.Email, user.Password), nil
}
