package main

const a = "Hello world"

type ID int

var (
	b bool
	c int
	d string
	f ID = 3
)

func main() {
	println(f)
}
