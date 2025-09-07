package application

import (
	"context"
	"fmt"

	"github.com/nick6969/go-clean-project/internal/config"
	"github.com/nick6969/go-clean-project/internal/logger"
)

type Application struct {
	Config *config.Config
	Logger *logger.Slogger

	Service *Service
	UseCase *UseCase
}

func New(cfg *config.Config) (*Application, error) {
	ctx := context.Background()
	logger := logger.NewSLogger(ctx, cfg.Logger)

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	service, err := NewService(app)
	if err != nil {
		return nil, fmt.Errorf("failed to create services: %w", err)
	}
	app.Service = service

	app.UseCase = NewUseCase(app)
	return app, nil
}
