package service

import (
	"context"
	"github.com/Peterliang233/ddd/go-task/domain/model/entity"
)

type ITaskService interface {
	CreateTask(ctx context.Context, task *entity.TaskEntity) (err error)
}
