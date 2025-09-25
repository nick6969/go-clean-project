package nxcache

import (
	"context"
	"log"
	"time"

	"github.com/nick6969/go-clean-project/internal/domain"
)

type job[T any] struct {
	// 用來標識這個請求
	LockIdentify string
	// 從快取讀取資料的函數
	CacheGetter func(ctx context.Context) (*T, error)
	// 將資料寫入快取的函數
	CacheSetter func(ctx context.Context, value *T) error
	// 從資料庫或其他來源一次性讀取資料的函數
	OnceGetter func(ctx context.Context) (*T, error)
	// lock 鎖住的最大時間，超過這個時間會自動釋放鎖
	LockTTL time.Duration
	// lock 鎖住時的等待重取時間
	RetryWaitTime time.Duration
}

func (job job[T]) doWithLock(ctx context.Context, lockManager domain.LockManager) (*T, error) {
	// 1. 先從快取取得資料
	models, err := job.CacheGetter(ctx)
	if err == nil {
		return models, nil
	}

	// 2. 嘗試獲取分散式鎖
	lock := lockManager.NewLock(job.LockIdentify, job.LockTTL)
	locked, err := lock.TryLock(ctx)
	if err != nil {
		return nil, err
	}
	if locked {
		defer func() {
			if err := lock.Unlock(ctx); err != nil {
				log.Printf("failed to unlock: %v", err)
			}
		}() // 確保在函數結束時釋放鎖
		// 3. 獲取鎖成功，從資料庫取得資料
		models, err = job.OnceGetter(ctx)
		if err != nil {
			return nil, err
		}

		// 4. 將資料寫入快取
		err = job.CacheSetter(ctx, models)
		if err != nil {
			return nil, err
		}

		// 5. 回傳資料
		return models, nil
	}

	// 6. 獲取鎖失敗，表示有其他請求正在更新快取
	// 可以選擇等待一段時間後重試，或者直接回傳錯誤
	// 這裏我們實作等待一段時間後重試，等待的時間可以根據實際情況調整
	time.Sleep(job.RetryWaitTime)
	return job.CacheGetter(ctx)
}
