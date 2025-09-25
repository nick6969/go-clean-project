package singlefight

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/domain"
	"golang.org/x/sync/singleflight"
)

type repo interface {
	GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error)
}

type cache interface {
	GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error)
	SetBranchInfos(ctx context.Context, value *[]domain.BranchInfo) error
}

type Service struct {
	repo  repo
	cache cache
	group group
}

func NewService(repo repo, cache cache) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
		group: &singleflight.Group{},
	}
}

func (s *Service) GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error) {
	job := job[[]domain.BranchInfo]{
		WorkIdentify: "GetBranchInfos",
		CacheGetter:  s.cache.GetBranchInfos,
		CacheSetter:  s.cache.SetBranchInfos,
		OnceGetter:   s.repo.GetBranchInfos,
	}

	return job.doWith(ctx, s.group)
}
