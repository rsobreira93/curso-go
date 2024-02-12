package main

import "github.com/rsobreira93/curso-go/fcutils/pkg/rabbitmq"

func main() {
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()

	rabbitmq.Publish(channel, "Hello World", "amq.direct")
}
