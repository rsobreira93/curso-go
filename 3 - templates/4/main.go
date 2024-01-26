package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("templates.html").ParseFiles("templates.html"))

		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 35},
			{"C#", 50},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
