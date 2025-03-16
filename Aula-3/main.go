package main
import 	"fmt"
// declaração de variáveis GLOBAIS, você pode acessar em todas as funções dentro do arquivos / pagote
// variáveis de escopo GLOBAIS

type ID int // Você pode cria seu próprio TYPE 

var (
	f ID = 1
)

func main() {

	println(f)

	fmt.Printf("O tipo de F é %T \n", f) // print do tipo da variável 
}