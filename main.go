package main

import (
	"encoding/json"
	"fmt"

	"github.com/estudo-udemy-go/slice"

	"github.com/estudo-udemy-go/estrutura"
	"github.com/estudo-udemy-go/funcao"
	"github.com/estudo-udemy-go/ponteiro"
	"github.com/estudo-udemy-go/variavel"
)

func main() {
	//testaVariavelPacotes()
	//testaFuncoes()
	//instanciandoEstruturas()
	//funcaoEstrutura()
	//convertePJson()
	//condicional.TestaIF()
	//loops.TestaLoop()
	//	condicional.TestaSwitch()
	//colecao.TestaMapa()
	//erro.TestaErros()
	//array.TestaArrays()
	slice.TestaSlice()
}

func funcaoEstrutura() {
	carro := estrutura.Carro{}
	fmt.Printf("Carro:  %+v\r\n", carro)
	carro.Modelo = "Crossfox"
	carro.Marca = "Volks"
	//carro.valor= Impossivel atribuir pois é privado
	carro.SetValor(50000.00)
	fmt.Printf("Carro apos set:  %+v\r\n", carro)
	fmt.Printf("Valor via get : %.2f\n", carro.GetValor())

}

func convertePJson() {
	carro := estrutura.Carro{}
	fmt.Printf("Carro:  %+v\r\n", carro)
	carro.Modelo = "Crossfox"
	carro.Marca = "Volks"
	objJSON, _ := json.Marshal(carro)
	fmt.Println("Carro em JSON" + string(objJSON))
}

func exemploPonteiro() {
	//Obtem o endereco de memoria da estrutura
	casa := new(estrutura.Imovel)

	fmt.Printf("CAsa: %p - %+v\r\n", &casa, casa)
	//Construtor com um parametro
	casa2 := estrutura.Imovel{Valor: 2000}

	fmt.Printf("CAsa: %p - %+v\r\n", &casa2, casa2)

	ponteiro.MudaDados(&casa2)

	fmt.Printf("CAsa2 apos mudarDados:  %+v\r\n", casa2)
}

func instanciandoEstruturas() {
	estrutura.CriacaoEstrutura()
}

//TestaVariavelPacotes executa chamadas de fora do pacote
func testaVariavelPacotes() {

	/**
	  LEMBRAR

	  Não é possivel no go ter pacotes com importacao ciclica

	  **/

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
