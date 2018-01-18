package variavel

import (
	"fmt"
	"strconv"
)

var (
	//Endereco é uma variavel publica, por isso comeca com maisculo
	Endereco    = "Rua X"
	telefone    string
	ddd, numero int
	privado     = 5
)

// ImprimeVariaveis TESTE DE METODO PUBLICO //
func ImprimeVariaveis() {

	testeInstanciaDireta := "Testandoo"
	fmt.Printf("endereco: %s\r\n", Endereco)

	fmt.Printf("variavel diretamente : %s\r\n", testeInstanciaDireta)

	fmt.Printf("variavel privada : %v\r\n", privado)

	fmt.Printf("variavel outro arquivo mesmo pacote : %v\r\n", dentroPacote)

	modelo = "Fox"

	fmt.Println("Atribuindo valor a variavel de outro arquivo do mesmo pacote e imprimindo : " + modelo)

	numero = 10

	//IMPRIMINDO ENDERECO FAZENDO A CONVERSAO DE NUMERO PARA STRING

	fmt.Println(Endereco + " numero " + strconv.Itoa(numero))

}
