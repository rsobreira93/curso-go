package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)

	go publish(channel)
	go reader(channel, &waitGroup)

	waitGroup.Wait()
}

func reader(ch chan int, waitGroup *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
		waitGroup.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
