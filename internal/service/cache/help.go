package cache

import (
	"context"
	"time"

	"github.com/nick6969/go-clean-project/pkg/gob"
	"github.com/nick6969/go-clean-project/pkg/json"
)

func getJSONModel[T any](ctx context.Context, key string, s *Service) (*T, error) {
	var value json.Container[T]
	err := s.client.GetModel(ctx, key, &value)
	if err != nil {
		return nil, err
	}
	return &value.RawValue, nil
}

func setJSONModel[T any](ctx context.Context, key string, model *T, expire time.Duration, s *Service) error {
	value := json.Container[T]{RawValue: *model}
	return s.client.SetModel(ctx, key, value, expire)
}

func setGOBModel[T any](ctx context.Context, key string, model *T, expire time.Duration, s *Service) error {
	value := gob.Container[T]{RawValue: *model}
	return s.client.SetModel(ctx, key, value, expire)
}

func getGOBModel[T any](ctx context.Context, key string, s *Service) (*T, error) {
	var value gob.Container[T]
	err := s.client.GetModel(ctx, key, &value)
	if err != nil {
		return nil, err
	}
	return &value.RawValue, nil
}
