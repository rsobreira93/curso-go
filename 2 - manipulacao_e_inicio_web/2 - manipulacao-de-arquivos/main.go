package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Olá mundo!")
	tamanho, err := f.Write([]byte("Olá mundo!"))

	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo criado com sucesso!\n Tamanho: ", tamanho, "bytes")

	f.Close()

	// Leitura do arquivo

	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo

	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			fmt.Println("Erro ao ler o arquivo: ", err)
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// defer f.Close()
}
