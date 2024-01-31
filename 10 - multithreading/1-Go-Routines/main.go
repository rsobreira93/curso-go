package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %v is running\n", i, name)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(time.Second * 1)
		}
	}()

	time.Sleep(time.Second * 15)
}
