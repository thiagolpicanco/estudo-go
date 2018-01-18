package estrutura

import (
	"fmt"
)

type Imovel struct {
	Valor  int
	Nome   string
	cidade string
}

func CriacaoEstrutura() {
	//Algo como um construtor vazio
	casa := Imovel{}
	fmt.Printf("Casa: %+v\r\n", casa)

	//Algo como um construtor por parametros
	apartamento := Imovel{50000, "Apzao", "Saqua"}
	fmt.Printf("AP: %+v\r\n", apartamento)

	//Algo como um constructor formado por um json
	sitio := Imovel{
		Valor:  100000,
		Nome:   "Sitio dona marta",
		cidade: "Arraial do cabo",
	}

	fmt.Printf("Sitio: %+v\r\n", sitio)

	//Atribuicao de valor
	sitio.cidade = "Araruama"

	fmt.Printf("Sitio: %+v\r\n", sitio)

}
