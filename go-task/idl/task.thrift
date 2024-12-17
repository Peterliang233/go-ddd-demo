namespace go example.task.item

include "base.thrift"

struct TaskItem {
    1: i64 id
    2: string name
    3: string description
    4: optional bool is_all
    5: TaskPeople target
}

struct TaskPeople {
    1: i64 target_type
    2: string group_code
    3: i64 cnt
}

struct GetTaskReq {
    1: required i64 id
}

struct GetTaskResp {
    1: TaskItem task

    255: base.BaseResp baseResp
}

service ItemService{
    GetTaskResp GetTask(1: GetTaskReq req)
}