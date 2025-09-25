package application

import (
	"context"
	"embed"
	"fmt"

	"github.com/nick6969/go-clean-project/docs/api"
	"github.com/nick6969/go-clean-project/internal/config"
	"github.com/nick6969/go-clean-project/internal/database/mysql"
	"github.com/nick6969/go-clean-project/internal/database/redis"
	"github.com/nick6969/go-clean-project/internal/logger"
)

type Application struct {
	Config   *config.Config
	Embed    Embeds
	Logger   *logger.Slogger
	Database *mysql.Database
	Redis    *redis.Client

	Service *Service
	UseCase *UseCase
}

func New(cfg *config.Config) (*Application, error) {
	ctx := context.Background()
	logger := logger.NewSLogger(ctx, cfg.Logger)
	database, err := mysql.InitDatabase(cfg.MySQL.DSN(), logger.GetDatabaseLogger())
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.NewClient(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}

	app := &Application{
		Config: cfg,
		Embed: Embeds{
			APIDoc: api.FS,
		},
		Logger:   logger,
		Database: database,
		Redis:    redisClient,
	}

	service, err := NewService(app)
	if err != nil {
		return nil, fmt.Errorf("failed to create services: %w", err)
	}
	app.Service = service

	app.UseCase = NewUseCase(app)
	return app, nil
}

type Embeds struct {
	APIDoc embed.FS
}
