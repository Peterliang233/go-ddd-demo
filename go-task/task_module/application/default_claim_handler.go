package application

import (
	"DDD/go-ddd-demo/go-task/kitex_gen/example/go/task"
	"context"
)

var TaskHandler ITaskHandler

type ITaskHandler interface {
	GetTask(ctx context.Context, req *task.GetTaskReq) (resp *task.GetTaskResp, err error)
	CreateTask(ctx context.Context, req *task.CreateTaskReq) (resp *task.CreateTaskResp, err error)
}

type DefaultTaskHandler struct{}

func (*DefaultTaskHandler) GetTask(ctx context.Context, req *task.GetTaskReq) (resp *task.GetTaskResp, err error) {
	return
}

func (*DefaultTaskHandler) CreateTask(ctx context.Context, req *task.CreateTaskReq) (resp *task.CreateTaskResp, err error) {
	return
}
