package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

// Método para desativar o cliente
func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

// Estrutura Empresa
type Empresa struct {
	Nome  string
	Ativa bool
}

// Método para desativar a empresa
func (e *Empresa) Desativar() {
	e.Ativa = false
	fmt.Printf("A empresa %s foi desativada\n", e.Nome)
}

// Interface Pessoa
type Pessoa interface {
	Desativar()
}

// Função que desativa qualquer tipo que implemente a interface Pessoa
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	Guilherme := Cliente{
		Nome:  "Guilherme",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua 1",
			Numero:     123,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}

	// Desativando o cliente usando a função Desativacao
	Desativacao(&Guilherme)

	fmt.Println("Status ativo do cliente:", Guilherme.Ativo)
	fmt.Println("Endereço do cliente:", Guilherme.Endereco)

	// Criando uma empresa e desativando-a
	minhaEmpresa := Empresa{
		Nome:  "TechCorp",
		Ativa: true,
	}

	Desativacao(&minhaEmpresa)
	fmt.Println("Status ativo da empresa:", minhaEmpresa.Ativa)
}
