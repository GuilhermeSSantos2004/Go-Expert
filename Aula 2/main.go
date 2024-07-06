package main

// declaração de variáveis GLOBAIS, você pode acessar em todas as funções dentro do arquivos / pagote
// variáveis de escopo GLOBAIS
var (
	a string
	b bool
	c int
	d float64
)

func main() {

	var e string = "Guilherme" // Declaração e atruibução no exato momento
	var f string               // variável de escopo local

	println(a) // imprimi em branco (nada) por padrão
	println(b) // imprimi um valor bool (FALSE) por padrão
	println(c) // imprimi um valor int (0) por padrão
	println(d) // imprimi um valor float64 ou float32 (0.000000.000) por padrão
	println(e) // imprimi a Declaração e atruibução de uma variavel que foi declarada no exato momento
	println(f) // imprimi um valor de uma variável de escopo local
}
