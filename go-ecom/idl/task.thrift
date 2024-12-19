namespace go example.go.ecom

include "base.thrift"

struct TaskItem {
    1: i64 id
    2: string name
    3: string description
    4: bool is_all
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

struct CreateTaskReq {
    1: string name
    2: string desc
    3: bool is_all
    4: TaskPeople people
}

struct CreateTaskResp {
    1: i64 task_id

    255: base.BaseResp baseResp
}

