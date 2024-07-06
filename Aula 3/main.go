package main

// declaração de variáveis GLOBAIS, você pode acessar em todas as funções dentro do arquivos / pagote
// variáveis de escopo GLOBAIS

type ID int

var (
	f ID = 1
)

func main() {

	println(f)

}
