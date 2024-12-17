package application

import (
	"DDD/go-ddd-demo/go-task/kitex_gen/example/task/item"
	"context"
)

var TaskHandler ITaskHandler

type ITaskHandler interface {
	GetTask(ctx context.Context, req *item.GetTaskReq) (resp *item.GetTaskResp, err error)
}

type DefaultTaskHandler struct{}

func (*DefaultTaskHandler) GetTask(ctx context.Context, req *item.GetTaskReq) (resp *item.GetTaskResp, err error) {
	return
}
