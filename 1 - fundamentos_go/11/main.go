package main

import "fmt"

type Address struct {
	street string
	number int
	city   string
	state  string
}

type Client struct {
	name   string
	age    int
	active bool
	Address
}

type Person interface {
	Desactive()
}

func (client Client) Desactive() {
	client.active = false
	fmt.Println(client)
}

func desativacao(pessoa Person) {
	pessoa.Desactive()
}

func main() {
	client := Client{
		name:   "John Doe",
		age:    18,
		active: true,
	}
	client.street = "Street 1"
	client.number = 123
	client.city = "City"

	client.Desactive()

}
