package main

import "fmt"

type Client struct {
	name string
}

func (c *Client) andou() {
	c.name = "João dos pao"
	fmt.Printf("O cliente %s andou", c.name)
}

func main() {

	client := Client{name: "João"}
	client.andou()

	fmt.Println("O client", client.name)
}
