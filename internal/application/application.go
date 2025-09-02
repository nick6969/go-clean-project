package application

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/config"
	"github.com/nick6969/go-clean-project/internal/logger"
)

type Application struct {
	Config *config.Config
	Logger logger.Logger
}

func New(cfg *config.Config) (*Application, error) {
	ctx := context.Background()
	logger := logger.NewSLogger(ctx, cfg.Logger)

	return &Application{
		Config: cfg,
		Logger: logger,
	}, nil
}
