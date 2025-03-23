package main

import "fmt"

func main() {
	var meuArray [3]int
	var seuArray [6]int
	var testArray[3]int

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	seuArray[0] = 1
	seuArray[1] = 2
	seuArray[2] = 3
	seuArray[3] = 4
	seuArray[4] = 5
	seuArray[5] = 6

	testArray[0] = 3
	testArray[1] = 6
	testArray[2] = 9

	fmt.Println(len(testArray) - 1) // O 'len' retona o TAMANHO do 'testeArray' - 1 
	fmt.Println(testArray[len(testArray) - 1])
	
	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor que está localizado no indice é: %d\n", i, v)
	}

	for i, v := range seuArray {
		fmt.Printf("O valor do indice é %d e o valor que está localizado no indice é: %d\n", i, v)
	}
}
