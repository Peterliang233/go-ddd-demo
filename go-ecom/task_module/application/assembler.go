package application

import (
	"DDD/go-ddd-demo/go-ecom/kitex_gen/example/go/ecom"
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/entity"
	"DDD/go-ddd-demo/go-ecom/task_module/domain/model/vo"
	"context"
	"errors"
)

func AssemblerCreate2Entity(ctx context.Context, req *ecom.CreateTaskReq) (task *entity.TaskEntity, err error) {
	if req == nil {
		return nil, errors.New("nil task request")
	}
	return &entity.TaskEntity{
		Name:      req.Name,
		Desc:      req.Desc,
		AllTarget: req.IsAll,
		Target:    vo.NewGroupPeople(req.People.TargetType, req.People.GroupCode, req.People.Cnt),
	}, nil
}

func AssemblerTaskEntity2AO(ctx context.Context, task *entity.TaskEntity) (resp *ecom.TaskItem, err error) {
	if task == nil {
		return nil, errors.New("nil task")
	}
	return &ecom.TaskItem{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Desc,
		IsAll:       task.AllTarget,
		Target: &ecom.TaskPeople{
			TargetType: task.Target.TargetType,
			GroupCode:  task.Target.GroupCode,
			Cnt:        task.Target.Cnt,
		},
	}, nil
}
