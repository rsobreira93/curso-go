package matematica

// Se a função ou variavel ou struct começa com letra maiuscula, ela é publica
// Se a função ou variavel ou struct começa com letra minuscula, ela é privada

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Andando..."
}
