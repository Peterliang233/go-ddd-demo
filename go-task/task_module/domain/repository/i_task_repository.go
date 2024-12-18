package repository

import (
	"DDD/go-ddd-demo/go-task/task_module/domain/model/entity"
	"context"
)

type ITaskRepository interface {
	CreateTask(ctx context.Context, task *entity.TaskEntity) error
}
