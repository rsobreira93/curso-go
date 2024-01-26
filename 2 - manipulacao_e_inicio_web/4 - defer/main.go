package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	defer request.Body.Close() // o defer é executado no final da função

	res, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))

	request.Body.Close()
}
