package singlefight

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"
)

// 定義一個 interface，包含 Do 和 DoChan 方法(幫 singleflight 做 interface)
type group interface {
	Do(key string, fn func() (any, error)) (v any, err error, shared bool)
	DoChan(key string, fn func() (any, error)) <-chan singleflight.Result
}

// 定義一個 job 結構體，包含工作標識、快取讀取函數、快取寫入函數、一次性讀取函數和預設值
type job[T any] struct {
	// 用來標識這個請求
	WorkIdentify string
	// 從快取讀取資料的函數
	CacheGetter func(ctx context.Context) (*T, error)
	// 將資料寫入快取的函數
	CacheSetter func(ctx context.Context, value *T) error
	// 從資料庫或其他來源一次性讀取資料的函數
	OnceGetter func(ctx context.Context) (*T, error)
}

// 定義一個方法，使用 singleflight 來執行工作
func (job job[T]) doWith(ctx context.Context, engine group) (*T, error) {
	// 使用 singleflight 的 Do 方法來執行工作
	// 這裏就保證同一時間只有一個請求會去執行
	m, err, _ := engine.Do(job.WorkIdentify, func() (any, error) {
		// 1. 先從快取取得資料
		v, e := job.CacheGetter(ctx)
		if e == nil {
			return v, nil
		}

		// 2. 如果快取沒有命中，則從資料庫或其他來源取得資料
		v, e = job.OnceGetter(ctx)
		if e != nil {
			return nil, e
		}

		// 3. 將資料寫入快取
		e = job.CacheSetter(ctx, v)
		if e != nil {
			return nil, e
		}

		// 4. 回傳資料
		return v, nil
	})

	if err != nil {
		return nil, err
	}

	value, ok := m.(*T)
	if !ok {
		return nil, errors.New("type assertion failed")
	}

	return value, nil
}
