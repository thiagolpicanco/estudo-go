package funcao

import (
	"fmt"
	"strconv"
)

func main() {
	resultado := Multiplicar(2, 3)
	fmt.Printf("O resultado foi : %d\r\n", resultado)

}

/*
Multiplicar faz a multiplicacao entre numeros
*/
func Multiplicar(x int, y int) int {
	fmt.Printf("Executando multiplicacao de %d com %d\n", x, y)
	return x * y

}

/*
Somar efetua a soma
*/
func Somar(x int, y int) int {
	fmt.Printf("Executando soma de %d com %d\n", x, y)
	return x + y
}

//ConverteIntParaString Converte num em string
func ConverteIntParaString(numero int) string {
	return strconv.Itoa(numero)
}
