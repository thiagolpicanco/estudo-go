package funcao

import (
	"fmt"
)

//DivisaoComRetornoAssinado efetua a divisao de numero1 com numero2 levando uma variavel ja assinada
func DivisaoComRetornoAssinado(numero1 int, numero2 int) (resultado int) {
	resultado = numero1 / numero2
	fmt.Printf("Efetuando divisao de %d com %d\r\n", numero1, numero2)
	fmt.Printf("Resultado %d\r\n", resultado)
	return
}

//RestoDaDivisao retorna o resto com variavel de retorno assinada
func RestoDaDivisao(numero1 int, numero2 int) (resto int) {
	fmt.Printf("Obtendo resto de %d com %d\r\n", numero1, numero2)

	resto = numero1 % numero2
	fmt.Printf("Resultado %d\r\n", resto)
	return
}

//DivisaoComResto retorna 2 valores assinados na mesma funcao
func DivisaoComResto(numero1, numero2 int) (divisao int, resto int) {
	divisao = DivisaoComRetornoAssinado(numero1, numero2)
	resto = RestoDaDivisao(numero1, numero2)
	return
}
