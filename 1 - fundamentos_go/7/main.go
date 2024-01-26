package main

import "fmt"

func main() {
	salarios := map[string]float64{"José João": 11325.45, "Gabriela Silva": 15456.78, "Pedro Junior": 1200.0}
	// fmt.Println(salarios["José João"])
	delete(salarios, "José João")
	// fmt.Println(salarios)
	salarios["Romulo"] = 5000.0
	// fmt.Println(salarios["Romulo"])

	// sal := make(map[string]float64)
	// sal1 := map[string]float64{}

	for nome, salario := range salarios {
		fmt.Println(nome, salario)
	}

	for _, salario := range salarios {
		fmt.Println(salario)
	}
}
