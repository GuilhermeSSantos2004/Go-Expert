package main 

string firstcode = "Hello World"

func main(){
	//O nome do pacote deve ser o nome da pasta onde está o arquivo.
	//Erros : Não pode ter dois pagotes dentro da mesma pasta
	// Curiosidade: As variáveis que são declaradas no pagote A podem 
	// ser vistas como variáveis globais em todos os arquivos que estiverem declarados o pacote XA mas não no pagote XB 
	print(firstcode)
}

