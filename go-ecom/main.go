package main

import (
	"DDD/go-ddd-demo/go-ecom/kitex_gen/example/go/ecom/ecomservice"
)

func main() {
	svr := ecomservice.NewServer(new(EcomServiceImpl))
	if err := svr.Run(); err != nil {
		panic(err)
	}
}
