package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var Producer rocketmq.Producer

var Consumer rocketmq.PushConsumer

func InitRMQ() {
	// init producer
	Producer, _ = rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2))

	if err := Producer.Start(); err != nil {
		panic(err)
	}

	// init consumer
	Consumer, _ = rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	if err := Consumer.Subscribe("min", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		fmt.Printf("subscribe callback: %v \n", msgs)
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		panic(err)
	}
	if err := Consumer.Start(); err != nil {
		panic(err)
	}
}
