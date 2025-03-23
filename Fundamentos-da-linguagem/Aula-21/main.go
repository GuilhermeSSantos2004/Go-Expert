package main

// =======================
// Como gerar binários em diferentes sistemas operacionais:
// =======================

// Para gerar um executável para Windows:
// GOOS=windows go build main.go

// Para gerar um executável para Linux:
// GOOS=linux go build main.go

// Para gerar um executável para macOS:
// GOOS=darwin go build main.go

// =======================
// Explicações sobre as variáveis de ambiente do Go:
// =======================

// GOOS => Define o sistema operacional de destino (ex: linux, windows, darwin)
// GOARCH => Define a arquitetura da CPU de destino (ex: amd64, arm64, 386)

// Para listar todas as combinações possíveis de GOOS e GOARCH:
// go tool dist list

// Para ver as configurações atuais da sua máquina:
// go env GOOS GOARCH

// Para gerar um binário com nome específico (ex: xpto):
// go build -o xpto main.go

func main() {
	a := 1

	switch a {
	case 1:
		print("O valor de 'a' é", a)
	case 2:
		print("O valor de 'a' é", a)
	case 3:
		print("O valor de 'a' é", a)
	default:
		print("Default: D")
	}
}
