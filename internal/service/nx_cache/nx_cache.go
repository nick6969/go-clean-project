package nxcache

import (
	"context"
	"time"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type repo interface {
	GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error)
}

type cache interface {
	GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error)
	SetBranchInfos(ctx context.Context, value *[]domain.BranchInfo) error
}

type Service struct {
	repo        repo
	cache       cache
	lockManager domain.LockManager
}

func NewService(repo repo, cache cache, lockManager domain.LockManager) *Service {
	return &Service{
		repo:        repo,
		cache:       cache,
		lockManager: lockManager,
	}
}

func (s *Service) GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error) {
	job := job[[]domain.BranchInfo]{
		LockIdentify:  "branch_infos_lock",
		CacheGetter:   s.cache.GetBranchInfos,
		CacheSetter:   s.cache.SetBranchInfos,
		OnceGetter:    s.repo.GetBranchInfos,
		LockTTL:       5 * time.Second,
		RetryWaitTime: 100 * time.Millisecond,
	}
	return job.doWithLock(ctx, s.lockManager)
}
