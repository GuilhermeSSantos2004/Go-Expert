package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(5, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

	sub, codicao := sub(0, 1)

	fmt.Printf("O valor retornado é %d a condição é %t", sub, codicao)
}

func sum(a, b int) (int, error) { //nome_func(var_receber <type>,var_receber <type>) (<type>, <type>) => Esse segundo parentes é para os tipos dos valores que serão retornados
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil //Essa func retorna um INTEIRO e um err => nil é nulo, vazio
}

func sub(a, b int) (int, bool) {

	if a-b <= 0 {
		return a - b, false
	}
	return a - b, true
}

func soma(a int, b int) (int) {  // indica o retorno da função se remover o '(int)' e a função tiver um retorno não irá funcionar 
	return a + b
}
