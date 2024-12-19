package repository_impl

import (
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/entity"
	"DDD/go-ddd-demo/go-ecom/task_module/infra/dependency/mysql/do"
)

func TaskEntity2Do(e *entity.TaskEntity) *do.TaskDo {
	return &do.TaskDo{
		Id:   e.Id,
		Name: e.Name,
		Desc: e.Desc,
	}
}

func TaskDo2Entity(d *do.TaskDo) *entity.TaskEntity {
	return &entity.TaskEntity{
		Id:   d.Id,
		Name: d.Name,
		Desc: d.Desc,
	}
}
