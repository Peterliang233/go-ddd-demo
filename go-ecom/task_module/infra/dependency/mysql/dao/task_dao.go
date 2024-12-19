package dao

import (
	"DDD/go-ddd-demo/go-ecom/common_module/infra/mysql"
	"DDD/go-ddd-demo/go-ecom/task_module/infra/dependency/mysql/do"
	"context"
	"errors"
	"gorm.io/gorm"
)

func GetTaskById(ctx context.Context, id int64) (*do.TaskDo, error) {
	var task do.TaskDo
	err := mysql.GormDB.WithContext(ctx).Debug().Where("id = ?", id).First(&task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func CreateTask(ctx context.Context, task *do.TaskDo) error {
	return mysql.GormDB.WithContext(ctx).Debug().Create(task).Error
}
