package main

import "fmt"

const a = "Hello world"

type ID int

var (
	b bool
	c int
	d string
	f ID = 3
)

func main() {
	fmt.Printf("O tipo de D é %T", d)
}
