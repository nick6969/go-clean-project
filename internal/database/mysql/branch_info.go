package mysql

import (
	"context"

	"github.com/nick6969/go-clean-project/internal/domain"
)

func (d *Database) GetBranchInfos(ctx context.Context) (*[]domain.BranchInfo, error) {
	// 這是說明用的假資料
	branchInfos := []domain.BranchInfo{
		{ID: 1, Name: "Main Branch", Address: "123 Main St"},
		{ID: 2, Name: "Secondary Branch", Address: "456 Side St"},
	}
	return &branchInfos, nil
}
