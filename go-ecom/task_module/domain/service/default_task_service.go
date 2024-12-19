package service

import (
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/entity"
	"DDD/go-ddd-demo/go-ecom/task_module/domain/repository"
	"context"
)

type DefaultTaskService struct{}

func (s *DefaultTaskService) CreateTask(ctx context.Context, task *entity.TaskEntity) (err error) {
	return repository.DefaultTaskRepository.CreateTask(ctx, task)
}

func (s *DefaultTaskService) GetTask(ctx context.Context, id int64) (task *entity.TaskEntity, err error) {
	task, err = repository.DefaultTaskRepository.GetTaskById(ctx, id)
	return
}
