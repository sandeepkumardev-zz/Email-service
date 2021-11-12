package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/adjust/rmq/v4"
)

var Queue rmq.Queue

func CreateQueue() {
	connection, err1 := rmq.OpenConnection(
		"my queue",
		"tcp",
		"localhost:6379",
		1,
		make(chan<- error),
	)
	if err1 != nil {
		panic(fmt.Sprintf("Error in establishing queue connection %s", err1.Error()))
	}

	defaultQueue, err2 := connection.OpenQueue("Default")
	if err2 != nil {
		panic("Error in creating queue")
	}

	defaultQueue.StartConsuming(10, time.Second)

	defaultQueue.AddConsumerFunc("email-consumer", SendEmail)

	Queue = defaultQueue
}

func Dispatch(payload interface{}) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic("Hey something is wrong in Dispatch func.")
	}

	Queue.PublishBytes(payloadBytes)
}
