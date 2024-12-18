package repository_impl

import (
	"DDD/go-ddd-demo/go-task/task_module/domain/model/entity"
	"context"
)

type DefaultTaskRepositoryImpl struct{}

func (*DefaultTaskRepositoryImpl) CreateTask(ctx context.Context, task *entity.TaskEntity) error {
	return nil
}
