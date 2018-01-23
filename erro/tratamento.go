package erro

import (
	"encoding/json"
	"fmt"

	"github.com/estudo-udemy-go/estrutura"
)

//TestaErros teste erro
func TestaErros() {

	carro := estrutura.Carro{Marca: "Fiat", Modelo: "Uno", Valor: 2000}
	fmt.Println(carro)

	objJSON, err := json.Marshal(carro)

	if err != nil {
		fmt.Println("Deu erro", err)
	}

	fmt.Println("Carro em Json", string(objJSON))

	//recebe o erro e ja trata ao mesmo tempo com a condicional
	if err := carro.SetValor(100); err != nil {
		fmt.Println("Erro aconteceu \n", err)

		if err == estrutura.ErrValorCarroAlto {
			fmt.Println("Erro tratado como se fosse um catch")
		}

	}

}
