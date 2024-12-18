package main

import (
	task "DDD/go-ddd-demo/go-task/kitex_gen/example/go/task"
	"DDD/go-ddd-demo/go-task/task_module/application"
	"context"
)

// TaskServiceImpl implements the last service interface defined in the IDL.
type TaskServiceImpl struct{}

// GetTask implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) GetTask(ctx context.Context, req *task.GetTaskReq) (resp *task.GetTaskResp, err error) {
	return application.TaskHandler.GetTask(ctx, req)
}

// CreateTask implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) CreateTask(ctx context.Context, req *task.CreateTaskReq) (resp *task.CreateTaskResp, err error) {
	return application.TaskHandler.CreateTask(ctx, req)
}
