package application

import (
	"DDD/go-ddd-demo/go-ecom/kitex_gen/example/go/ecom"
	"DDD/go-ddd-demo/go-ecom/task_module/domain/service"
	"context"
	"log"
)

var TaskHandler ITaskHandler

type ITaskHandler interface {
	GetTask(ctx context.Context, req *ecom.GetTaskReq) (resp *ecom.GetTaskResp, err error)
	CreateTask(ctx context.Context, req *ecom.CreateTaskReq) (resp *ecom.CreateTaskResp, err error)
}

type DefaultTaskHandler struct{}

func (*DefaultTaskHandler) GetTask(ctx context.Context, req *ecom.GetTaskReq) (resp *ecom.GetTaskResp, err error) {
	resp = &ecom.GetTaskResp{}
	if req == nil {
		log.Println("req is nil")
		return
	}
	taskEntity, err := service.TaskService.GetTask(ctx, req.Id)
	if err != nil {
		return resp, err
	}

	taskAo, err := AssemblerTaskEntity2AO(ctx, taskEntity)
	if err != nil {
		log.Println(err)
		return
	}

	resp.Task = taskAo

	return resp, nil
}

func (*DefaultTaskHandler) CreateTask(ctx context.Context, req *ecom.CreateTaskReq) (resp *ecom.CreateTaskResp, err error) {
	taskEntity, err := AssemblerCreate2Entity(ctx, req)
	if err != nil {
		log.Printf("AssemblerCreate2Entity err: %v\n", err)
		return
	}

	err = service.TaskService.CreateTask(ctx, taskEntity)
	if err != nil {
		log.Printf("service.TaskService err: %v\n", err)
		return
	}
	return
}
