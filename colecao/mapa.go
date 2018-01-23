package colecao

import (
	"fmt"

	"github.com/estudo-udemy-go/estrutura"
)

var cache map[string]estrutura.Carro

//TestaMapa testa mapa
func TestaMapa() {

	cache = make(map[string]estrutura.Carro)

	fox := estrutura.Carro{}
	fox.Modelo = "Fox"
	fox.Marca = "Volkswagen"
	fox.SetValor(50000)

	crossFox := estrutura.Carro{}
	crossFox.Modelo = "CrossFox"
	crossFox.Marca = "Volkswagen"
	crossFox.SetValor(80000)

	cache["Foxzao"] = fox

	fmt.Println("Existe foxZao no cache?")

	cache[crossFox.Modelo] = crossFox

	carrin, achou := cache["Foxzao"]
	if achou {
		fmt.Println("achamos ele", carrin.Modelo)
	}

	for chave, carro := range cache {
		fmt.Println(chave, carro)
	}

	delete(cache, crossFox.Modelo)

	fmt.Printf("O Cache possui %d item(s) \n", len(cache))

	fmt.Println(cache["Foxzao"])

}
