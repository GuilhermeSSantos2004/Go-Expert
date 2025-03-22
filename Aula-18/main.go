package main

import (
	"fmt"
	"github.com/GuilhermeSSantos2004/Go-Expert/Aula-18/matematica"
	"github.com/google/uuid"
)


func main() {
	s:= matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)//
	fmt.Println(uuid.New())
}
