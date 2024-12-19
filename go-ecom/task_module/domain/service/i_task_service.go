package service

import (
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/entity"
	"context"
)

var TaskService ITaskService

type ITaskService interface {
	CreateTask(ctx context.Context, task *entity.TaskEntity) (err error)
	GetTask(ctx context.Context, id int64) (task *entity.TaskEntity, err error)
}
