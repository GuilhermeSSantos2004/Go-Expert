package main

import "fmt"

func main() {
	var meuArray [3]int

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor que está localizado no indice é: %d\n", i, v)
	}
}
