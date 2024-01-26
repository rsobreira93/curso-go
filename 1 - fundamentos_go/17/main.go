package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func soma[T Number](m map[string]T) T {
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}

	m3 := map[string]MyNumber{"a": 1, "b": 2, "c": 3}

	var result = soma(m)
	var result2 = soma(m2)
	var result3 = soma(m3)

	println(result)
	println(result2)
	println(result3)
	println(Compara(10, 10.0))

}
