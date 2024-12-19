package repository

import (
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/entity"
	"context"
)

var DefaultTaskRepository ITaskRepository

type ITaskRepository interface {
	CreateTask(ctx context.Context, task *entity.TaskEntity) error
	GetTaskById(ctx context.Context, id int64) (*entity.TaskEntity, error)
}
