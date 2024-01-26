package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	curso := Curso{"Curso de Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Nome: {{.Nome}}\nCarga Horária: {{.CargaHoraria}} horas\n"))

	err := t.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}

}
