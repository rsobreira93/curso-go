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
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for i, v := range meuArray {
		fmt.Printf("O valor de i é %d e o valor de v é %d\n", i, v)
	}
}
