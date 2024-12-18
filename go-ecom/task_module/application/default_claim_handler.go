package application

import (
	"DDD/go-ddd-demo/go-ecom/kitex_gen/example/go/ecom"
	"context"
)

var TaskHandler ITaskHandler

type ITaskHandler interface {
	GetTask(ctx context.Context, req *ecom.GetTaskReq) (resp *ecom.GetTaskResp, err error)
	CreateTask(ctx context.Context, req *ecom.CreateTaskReq) (resp *ecom.CreateTaskResp, err error)
}

type DefaultTaskHandler struct{}

func (*DefaultTaskHandler) GetTask(ctx context.Context, req *ecom.GetTaskReq) (resp *ecom.GetTaskResp, err error) {
	return
}

func (*DefaultTaskHandler) CreateTask(ctx context.Context, req *ecom.CreateTaskReq) (resp *ecom.CreateTaskResp, err error) {
	return
}
