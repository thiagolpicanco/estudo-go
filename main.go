package main

import (
	"fmt"

	"github.com/estudo-udemy-go/funcao"

	"github.com/estudo-udemy-go/variavel"
)

func main() {
	//testaVariavelPacotes()
	testaFuncoes()
}

/**
LEMBRAR

Não é possivel no go ter pacotes com importacao ciclica

**/

//TestaVariavelPacotes executa chamadas de fora do pacote
func testaVariavelPacotes() {
	fmt.Println("Imprimindo valor de variavel dentro de outro pacote por ser publica: " + variavel.Endereco)
	variavel.ImprimeVariaveis()
}

func testaFuncoes() {

	numero1 := 2
	numero2 := 3

	resultadoMulti := funcao.Multiplicar(numero1, numero2)
	fmt.Println(funcao.ConverteIntParaString(resultadoMulti))

	//Recebe a funcao soma como valor
	soma := funcao.Somar

	//Teste de funcao estilo callback

	//Envia a funcao soma como parametro, assim como os 2 numeros para serem calculado
	resultadoSoma := funcao.Calcular(soma, numero1, numero2)
	fmt.Println(resultadoSoma)

	numero1 = 101
	numero2 = 10

	resultadoDivisao := funcao.DivisaoComRetornoAssinado(numero1, numero2)
	fmt.Println(resultadoDivisao)

	resultadoDivisao2, resto := funcao.DivisaoComResto(numero1, numero2)

	fmt.Printf("Resultado da divisao: %d, Resto da divisao : %d\r\n", resultadoDivisao2, resto)

}
