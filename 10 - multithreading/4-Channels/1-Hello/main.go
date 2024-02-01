package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Hello" //jogando valor no canal
	}()

	message := <-channel //recebendo valor do canal
	fmt.Println(message)
}
