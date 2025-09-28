package redis

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisLock struct {
	client *redis.Client
	key    string
	value  string
	ttl    time.Duration
}

func NewRedisLock(client *redis.Client, key string, ttl time.Duration) *RedisLock {
	return &RedisLock{
		client: client,
		key:    key,
		value:  uuid.New().String(), // 使用 UUID 作為鎖的值，確保唯一性
		ttl:    ttl,
	}
}

func (l *RedisLock) TryLock(ctx context.Context) (bool, error) {
	// 使用 SET 命令設置鎖，NX 表示只有在鍵不存在時才設置，PX 表示設置鍵的過期時間
	result, err := l.client.SetNX(ctx, l.key, l.value, l.ttl).Result()
	if err != nil {
		return false, err
	}
	return result, nil // 返回是否成功獲取鎖
}

func (l *RedisLock) Unlock(ctx context.Context) error {
	// 使用 Lua 腳本來確保只有持有鎖的客戶端才能釋放鎖
	script := `
    if redis.call("get", KEYS[1]) == ARGV[1] then
      return redis.call("del", KEYS[1])
    else
      return 0
    end
  `
	result, err := l.client.Eval(ctx, script, []string{l.key}, l.value).Result()
	if err != nil {
		return err
	}
	if result.(int64) == 0 { //nolint: errcheck // result 一定是 int64
		return errors.New("unlock failed: not the lock owner")
	}
	return nil
}
