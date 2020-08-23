package service

import "context"

// Service 服务层
type Service struct {
	ctx context.Context
}

// New 创建service
func New(ctx context.Context) Service {
	return Service{
		ctx: ctx,
	}
}
