package main

import "fmt"

func main() {
	// ponteiro , tamanho, capacidade (Slice '=?' Array)
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100} // Slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])  // :{indicador/p-visulizar} // : => ponto de corte slice do slice

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
}

