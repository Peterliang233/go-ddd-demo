package service

import (
	"DDD/go-ddd-demo/go-task/task_module/domain/model/entity"
	"context"
)

type DefaultTaskService struct{}

func (s *DefaultTaskService) CreateTask(ctx context.Context, task *entity.TaskEntity) (err error) {
	return
}
