package mysql

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const maxOpenConns = 25
const maxIdleConns = 3
const maxLiftTime = 5 * time.Minute

type Database struct {
	db *gorm.DB
}

func InitDatabase(dsn string, logger *slog.Logger) (*Database, error) {
	database, err := connectDB(dsn, maxOpenConns, maxIdleConns, maxLiftTime, logger)

	if err != nil {
		return nil, err
	}

	return &Database{database}, nil
}

func connectDB(dsn string, maxOpenConns, maxIdleConns int, maxLiftTime time.Duration, slogger *slog.Logger) (*gorm.DB, error) {
	logger := logger.NewSlogLogger(
		slogger,
		logger.Config{LogLevel: logger.Info, ParameterizedQueries: false, IgnoreRecordNotFoundError: true},
	)
	config := gorm.Config{
		Logger:                 logger,
		SkipDefaultTransaction: true,
	}

	database, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		return nil, err
	}

	sql, err := database.DB()
	if err != nil {
		return nil, err
	}

	sql.SetMaxOpenConns(maxOpenConns)
	sql.SetMaxIdleConns(maxIdleConns)
	sql.SetConnMaxLifetime(maxLiftTime)

	return database, nil
}

func (d *Database) CheckEmailIsExists(ctx context.Context, email string) (bool, error) {
	// TODO: - implement this function
	return false, errors.New("not implemented")
}

func (d *Database) CreateUser(ctx context.Context, email, hashedPassword string) (int, error) {
	// TODO: - implement this function
	return 0, errors.New("not implemented")
}
