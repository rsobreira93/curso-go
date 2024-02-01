package main

func main() {
	channel := make(chan int, 5)

	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	channel <- 5

	println(<-channel)
}
