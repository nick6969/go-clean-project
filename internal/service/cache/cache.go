package cache

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/database/redis"
	"github.com/nick6969/go-clean-project/internal/domain"
)

type Service struct {
	client *redis.Client
}

func NewService(client *redis.Client) *Service {
	return &Service{client: client}
}

const expirationBranchInfosKey = 24 * 60 * 60 // 1 day

func (s *Service) getBranchInfosKey() string {
	return "branch_infos"
}

func (s *Service) GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error) {
	key := s.getBranchInfosKey()
	return getGOBModel[[]domain.BranchInfo](ctx, key, s)
}

func (s *Service) SetBranchInfos(ctx context.Context, model *[]domain.BranchInfo) error {
	key := s.getBranchInfosKey()
	return setGOBModel(ctx, key, model, expirationBranchInfosKey, s)
}

func (s *Service) VoidBranchInfos(ctx context.Context) error {
	key := s.getBranchInfosKey()
	return s.client.Del(ctx, key)
}
