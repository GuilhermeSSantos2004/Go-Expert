package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000

	// sal := make(map[string]int) prepara o map com o make 
	// sal1 := map[string]int{} preprar o map sem nenhum valor 
	// sal1["Wesley"] = 1000  adiciona uma chave e valor no map

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}
}