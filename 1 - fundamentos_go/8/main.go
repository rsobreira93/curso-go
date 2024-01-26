package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := soma(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(valor)
}

func soma(a, b int) (int, error) {
	if a+b > 10 {
		return 0, errors.New("O valor é maior que 10")
	}
	return a + b, nil
}
