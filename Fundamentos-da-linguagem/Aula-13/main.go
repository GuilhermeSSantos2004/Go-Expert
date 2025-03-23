package main


// Se você quer que os valores sejam modificados em diferentes partes do código, use ponteiros.
// Caso contrário, passe os valores diretamente.

func soma(a, b *int) int {
	*a = 50 // Modifica o valor original da variável apontada por 'a'(direto no endereço da memória)
	*b = 50 // Modifica o valor original da variável apontada por 'b'(direto no endereço da memória)
	return *a + *b
}

func divisao(c, e int) int {
	c = 20
	e = 5
	return c / e 
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	minhaVar1 -= 10 // Apenas modifica o valor da variável existente
	println("Resultado da soma:", soma(&minhaVar1, &minhaVar2))
	println("Novo valor de &minhaVar1:", minhaVar1) 
	println("Novo valor de &minhaVar2:", minhaVar2) 

	valor1 := 10 
	valor2 := 2
	// Chamando a função divisao e passando os valores das variáveis (não ponteiros)
	// Como Go passa os valores por cópia, valor1 e valor2 não serão modificados

	println("Valor enviado da variável 'valor1' == ", valor1, 
		"Valor enviado da variável 'valor2' == ", valor2, 
		"Resultado da divisão:", divisao(valor1, valor2))

	// Exibindo os valores após a chamada da função (eles permanecem inalterados)
	println("Valor da variável 'valor1':", valor1) 
	println("Valor da variável 'valor2':", valor2) 
}
