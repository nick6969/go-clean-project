package application

import (
	"github.com/nick6969/go-clean-project/internal/domain"
	"github.com/nick6969/go-clean-project/internal/listener"
	"github.com/nick6969/go-clean-project/internal/service/cache"
	"github.com/nick6969/go-clean-project/internal/service/dispatcher"
	"github.com/nick6969/go-clean-project/internal/service/email"
	nxcache "github.com/nick6969/go-clean-project/internal/service/nx_cache"
	"github.com/nick6969/go-clean-project/internal/service/password"
	"github.com/nick6969/go-clean-project/internal/service/sfnx"
	"github.com/nick6969/go-clean-project/internal/service/singlefight"
	"github.com/nick6969/go-clean-project/internal/service/token"
)

type Service struct {
	Password     *password.Service
	Token        *token.Service
	SingleFlight *singlefight.Service
	Cache        *cache.Service
	NxCache      *nxcache.Service
	Sfnx         *sfnx.Service
	Email        *email.Service
	Dispatch     *dispatcher.Service
}

func NewService(app *Application) (*Service, error) {
	passwordService := password.NewService()
	tokenService, err := token.NewService([]byte(app.Config.Token.Secret))
	if err != nil {
		return nil, err
	}

	cacheService := cache.NewService(app.Redis)

	singleFightService := singlefight.NewService(app.Database, cacheService)

	nxCacheService := nxcache.NewService(app.Database, cacheService, app.Redis)

	sfnxService := sfnx.NewService(nxCacheService)

	emailService := email.NewService()

	dispatcherService := dispatcher.NewService(app.Logger)

	dispatcherService.RegisterListener(domain.EventUserRegistered, listener.NewWelcomeEmail(emailService))

	return &Service{
		Password:     passwordService,
		Token:        tokenService,
		SingleFlight: singleFightService,
		Cache:        cacheService,
		NxCache:      nxCacheService,
		Sfnx:         sfnxService,
		Email:        emailService,
		Dispatch:     dispatcherService,
	}, nil
}
