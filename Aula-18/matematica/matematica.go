package matematica

func Soma[T int | float64](a, b T) T { //Todo que começa com letra macuscula estaá exportando para usar em ooutos aruqivos e tudo que não estiver não estar.
	return a + b
}

var A = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	 return "Carro andando"
}
