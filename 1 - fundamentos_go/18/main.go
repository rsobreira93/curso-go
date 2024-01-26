package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rsobreira93/curso-go/18/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)

	fmt.Println("O resultado da soma é: ", soma)

	fmt.Println(uuid.New())
}
