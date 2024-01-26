package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"n"` //tag - mapeamento de campos para json, como isso será apresentado no JSON
	Saldo  float64 `json:"s"` // se tiver - ele desconsidera esse campo
}

func main() {
	conta := Conta{Numero: 123, Saldo: 100.50}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n":123,"s":100.5}`)
	var contax Conta
	err = json.Unmarshal(jsonPuro, &contax)

	if err != nil {
		panic(err)
	}

	fmt.Println(contax)

}
