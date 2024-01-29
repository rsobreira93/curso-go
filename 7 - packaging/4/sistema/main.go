package main

import (
	"github.com/google/uuid"
	"github.com/rsobreira93/goexpert/7/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	print(uuid.New().String())
}
