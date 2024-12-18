namespace go example.go.ecom

include "task.thrift"


service EcomService {
    task.GetTaskResp GetTask(1: task.GetTaskReq req)
    task.CreateTaskResp CreateTask(1: task.CreateTaskReq req)
}