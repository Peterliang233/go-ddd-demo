package repository_impl

import (
	"context"
	"github.com/Peterliang233/ddd/go-task/domain/model/entity"
)

type DefaultTaskRepositoryImpl struct{}

func (*DefaultTaskRepositoryImpl) CreateTask(ctx context.Context, task *entity.TaskEntity) error {
	return nil
}
