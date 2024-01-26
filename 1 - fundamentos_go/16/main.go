package main

import "fmt"

func main() {
	var minhaVar interface{} = "ROmulo Lindo"
	println(minhaVar.(string))

	res, ok := minhaVar.(int)

	fmt.Printf("%v %v\n", res, ok)
}
