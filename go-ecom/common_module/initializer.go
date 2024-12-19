package common_module

import (
	"DDD/go-ddd-demo/go-ecom/common_module/infra/mq"
	"DDD/go-ddd-demo/go-ecom/common_module/infra/mysql"
	"DDD/go-ddd-demo/go-ecom/common_module/infra/redis"
)

func init() {
	mysql.InitGorm()
	mq.InitRMQ()
	redis.InitRedisClient()
}
