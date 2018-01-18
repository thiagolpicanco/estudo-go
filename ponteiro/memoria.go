package ponteiro

import (
	"fmt"

	"github.com/estudo-udemy-go/estrutura"
)

func main() {

	casa := new(estrutura.Imovel)

	fmt.Printf("CAsa: %p - %+v\r\n", &casa, casa)
	casa2 := estrutura.Imovel{Valor: 2000}

	fmt.Printf("CAsa: %p - %+v\r\n", &casa2, casa2)

	MudaDados(&casa2)

	fmt.Printf("CAsa2 apos mudarDados:  %+v\r\n", casa2)

}

// MudaDados altera o valor de campos sem ter que retornar a estrutura enviada
//Lembra um set(java)
func MudaDados(imovel *estrutura.Imovel) {
	imovel.Valor = 5000
}
