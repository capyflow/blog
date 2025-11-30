package service

import (
	"context"

	"github.com/capyflow/blog/model"
)

type SystemService struct {
	ctx   context.Context
	group string
}

// NewSystemService 创建系统服务
func NewSystemService(ctx context.Context, group string) *SystemService {
	return &SystemService{
		ctx:   ctx,
		group: group,
	}
}

// 更新AI模型配置
func (s *SystemService) UpsertAiConfig(ctx context.Context, aiConfig *model.AiConfig) error {
	return nil
}
