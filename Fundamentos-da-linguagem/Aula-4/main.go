package main

import (
	"fmt"
)

type ID int

var (
	f    ID
	nome string
)

func main() {

	fmt.Printf("O tipo do (f) É: %T", f)
	fmt.Printf(" | ")
	fmt.Printf("O tipo do (f) É: %T", nome)
}
