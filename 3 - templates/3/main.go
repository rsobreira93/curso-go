package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	t := template.Must(template.New("templates.html").ParseFiles("templates.html"))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 35},
		{"C#", 50},
	})

	if err != nil {
		panic(err)
	}

}
