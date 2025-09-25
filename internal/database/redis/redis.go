package redis

import (
	"context"
	"time"

	"github.com/nick6969/go-clean-project/internal/config"
	"github.com/nick6969/go-clean-project/internal/domain"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
}

func NewClient(ctx context.Context, config config.RedisConfig) (*Client, error) {
	const dbIndex = 1
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address(),
		Password: config.Password,
		DB:       dbIndex,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Client{
		client: rdb,
	}, nil
}

func (c *Client) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Client) SetModel(ctx context.Context, key string, value any, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *Client) GetModel(ctx context.Context, key string, value any) error {
	v := c.client.Get(ctx, key)
	if v.Err() != nil {
		return v.Err()
	}

	err := v.Scan(value)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Del(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

func (c *Client) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

func (c *Client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

func (c *Client) NewLock(key string, ttl time.Duration) domain.Lock {
	return NewRedisLock(c.client, key, ttl)
}
