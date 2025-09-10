package entity

import (
	"time"
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
