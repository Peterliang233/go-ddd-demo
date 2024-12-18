package task_module

import "DDD/go-ddd-demo/go-ecom/task_module/application"

func init() {
	application.TaskHandler = &application.DefaultTaskHandler{}
}
