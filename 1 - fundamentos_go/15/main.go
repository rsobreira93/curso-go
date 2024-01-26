package main

import "fmt"

func main() {
	var x interface{} = 7
	var y interface{} = "string"

	f(x)
	f(y)

}

func f(t interface{}) {
	fmt.Printf("O tipo de t é %T e o valor %v\n", t, t)
}
