package main

import "fmt"

func main() {
	// Memória -> Endereço -> Valor

	// Variável -> Armazena um valor e tem um endereço na memória
	// Ponteiro -> Armazena o endereço de uma variável

	a := 10
	var ponteiro *int = &a // O 'ponteiro' armazena o endereço de 'a'

	*ponteiro = 20 // Altera o valor de 'a' através do ponteiro

	b := &a // 'b' recebe o mesmo endereço de 'a', apontando para o mesmo local na memória

	*ponteiro = 40

	// Exibindo os valores e endereços de memória
	fmt.Println("Valor de a:", a)        
	fmt.Println("Endereço de a:", &a)    // Endereço na memória
	fmt.Println("Endereço de b:", &b)    // Endereço de b (que armazena o ponteiro de 'a')
	fmt.Println("Valor apontado por b:", *b) // (valor armazenado no endereço apontado por b)
	fmt.Println("Endereço armazenado em ponteiro:", ponteiro)
}
