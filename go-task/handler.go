package main

import (
	"DDD/go-ddd-demo/go-task/kitex_gen/example/task/item"
	"DDD/go-ddd-demo/go-task/task_module/application"
	"context"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetTask implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetTask(ctx context.Context, req *item.GetTaskReq) (resp *item.GetTaskResp, err error) {
	return application.TaskHandler.GetTask(ctx, req)
}
