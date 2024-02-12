package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {

	connection, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return channel, nil
}

func Consume(channel *amqp.Channel, out chan<- amqp.Delivery) error {
	messages, error := channel.Consume(
		"minhafila",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if error != nil {
		return error
	}

	for message := range messages {
		out <- message
	}

	return nil
}
