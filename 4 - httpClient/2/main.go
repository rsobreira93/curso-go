package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"nome":"Romulo"`))

	response, err := c.Post("http://www.google.com.br", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.CopyBuffer(os.Stdout, response.Body, nil)
}
