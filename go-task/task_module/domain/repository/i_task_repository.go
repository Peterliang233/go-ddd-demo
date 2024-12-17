package repository

import (
	"context"
	"github.com/Peterliang233/ddd/go-task/domain/model/entity"
)

type ITaskRepository interface {
	CreateTask(ctx context.Context, task *entity.TaskEntity) error
}
