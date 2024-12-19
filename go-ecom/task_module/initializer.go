package task_module

import (
	"DDD/go-ddd-demo/go-ecom/task_module/application"
	"DDD/go-ddd-demo/go-ecom/task_module/domain/repository"
	"DDD/go-ddd-demo/go-ecom/task_module/domain/service"
	"DDD/go-ddd-demo/go-ecom/task_module/infra/repository_impl"
)

func init() {
	application.TaskHandler = &application.DefaultTaskHandler{}
	service.TaskService = &service.DefaultTaskService{}
	repository.DefaultTaskRepository = &repository_impl.DefaultTaskRepositoryImpl{}
}
