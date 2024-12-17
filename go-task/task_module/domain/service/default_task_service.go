package service

import (
	"context"
	"github.com/Peterliang233/ddd/go-task/domain/model/entity"
)

type DefaultTaskService struct{}

func (s *DefaultTaskService) CreateTask(ctx context.Context, task *entity.TaskEntity) (err error) {
	return
}
