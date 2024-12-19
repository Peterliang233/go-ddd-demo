package main

import (
	"DDD/go-ddd-demo/go-ecom/kitex_gen/example/go/ecom"
	"DDD/go-ddd-demo/go-ecom/task_module/application"
	"context"
)

// EcomServiceImpl implements the last service interface defined in the IDL.
type EcomServiceImpl struct{}

// GetTask implements the EcomServiceImpl interface.
func (s *EcomServiceImpl) GetTask(ctx context.Context, req *ecom.GetTaskReq) (resp *ecom.GetTaskResp, err error) {
	return application.TaskHandler.GetTask(ctx, req)
}

// CreateTask implements the EcomServiceImpl interface.
func (s *EcomServiceImpl) CreateTask(ctx context.Context, req *ecom.CreateTaskReq) (resp *ecom.CreateTaskResp, err error) {
	return application.TaskHandler.CreateTask(ctx, req)
}
