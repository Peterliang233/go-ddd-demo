package main

import "DDD/go-ddd-demo/go-task/kitex_gen/example/go/task/taskservice"

func main() {
	svr := taskservice.NewServer(new(TaskServiceImpl))
	if err := svr.Run(); err != nil {
		panic(err)
	}
}
