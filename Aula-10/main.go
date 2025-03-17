package main

// Closures => Funções anônimas

import "fmt"

func main() {

	totalvezesdobro := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543) * 2
	}()
	fmt.Println(totalvezesdobro)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
