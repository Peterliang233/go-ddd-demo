package repository_impl

import (
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/entity"
	"DDD/go-ddd-demo/go-ecom/task_module/infra/dependency/mysql/dao"
	"context"
)

type DefaultTaskRepositoryImpl struct{}

func (*DefaultTaskRepositoryImpl) CreateTask(ctx context.Context, task *entity.TaskEntity) error {
	taskDo := TaskEntity2Do(task)
	return dao.CreateTask(ctx, taskDo)
}

func (*DefaultTaskRepositoryImpl) GetTaskById(ctx context.Context, id int64) (*entity.TaskEntity, error) {
	taskDo, err := dao.GetTaskById(ctx, id)
	if err != nil {
		return nil, err
	}
	if taskDo == nil {
		return nil, nil
	}
	return TaskDo2Entity(taskDo), nil
}
