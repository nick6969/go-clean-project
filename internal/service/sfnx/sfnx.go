package sfnx

import (
	"context"
	"errors"

	"github.com/nick6969/go-clean-project/internal/domain"
	"golang.org/x/sync/singleflight"
)

type group interface {
	Do(key string, fn func() (any, error)) (v any, err error, shared bool)
}

type nxCache interface {
	GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error)
}

type Service struct {
	nxCache nxCache
	group   group
}

func NewService(nxCache nxCache) *Service {
	return &Service{
		nxCache: nxCache,
		group:   &singleflight.Group{},
	}
}

func (s *Service) GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error) {
	v, err, _ := s.group.Do("GetBranchInfos:SFNX", func() (any, error) {
		return s.nxCache.GetBranchInfos(ctx)
	})
	if err != nil {
		return nil, err
	}

	value, ok := v.(*[]domain.BranchInfo)
	if !ok {
		return nil, errors.New("type assertion to *[]domain.BranchInfo failed")
	}

	return value, nil
}
