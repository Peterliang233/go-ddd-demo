package entity

import (
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/vo"
)

type TaskEntity struct {
	Id        int64
	Name      string
	Desc      string
	AllTarget bool
	Target    vo.GroupPeople
}
