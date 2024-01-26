package main

import "fmt"

func main() {
	a := 10
	var pointer *int = &a
	*pointer = 20

	b := &a
	*b = 30
	fmt.Println(a)
}
