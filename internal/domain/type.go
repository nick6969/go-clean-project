package domain

import (
	"context"
	"time"
)

type LockManager interface {
	NewLock(key string, ttl time.Duration) Lock
}

type Lock interface {
	TryLock(ctx context.Context) (bool, error)
	Unlock(ctx context.Context) error
}
