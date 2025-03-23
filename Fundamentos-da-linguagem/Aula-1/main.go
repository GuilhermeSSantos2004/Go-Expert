package main // package main é obrigatório para criar um programa executável

// Quando você usa package main, está dizendo ao compilador:
// "Este é o ponto de entrada do meu programa."

// O Go vai procurar uma função main() dentro do package main.
// É essa função que será executada quando você rodar o programa.

func main() {
	// O nome do pacote deve ser o nome da pasta onde está o arquivo.
	// Erros: Não pode ter dois pacotes dentro da mesma pasta
	// Curiosidade: As variáveis que são declaradas no pacote A podem 
	// ser vistas como variáveis globais em todos os arquivos que estiverem declarados no pacote A, 
	// mas não no pacote B.

	firstcode := "Hello World"
	print(firstcode)
}
