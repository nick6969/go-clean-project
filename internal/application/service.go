package application

import (
	"github.com/nick6969/go-clean-project/internal/service/cache"
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

	return &Service{
		Password:     passwordService,
		Token:        tokenService,
		SingleFlight: singleFightService,
		Cache:        cacheService,
		NxCache:      nxCacheService,
		Sfnx:         sfnxService,
	}, nil
}
