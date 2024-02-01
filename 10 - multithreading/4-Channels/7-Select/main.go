package main

import (
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{Id: 1, Msg: "Hello from RabbitMQ"}

			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{Id: 1, Msg: "Hello from Kafka"}

			c1 <- msg
		}
	}()

	for {
		select {
		case mg1 := <-c1:
			println("c1", mg1.Msg)

		case mg2 := <-c2:
			println("c2", mg2.Msg)

		case <-time.After(3 * time.Second):
			println("timeout")

			// default:
			// 	println("default")
		}
	}
}
