package funcao

//Calcular recebe uma funcao que tem parametros inteiros e retorna inteiro, assim como recebe 2 numeros para a operacao
func Calcular(funcao func(int, int) int, numero1 int, numero2 int) int {
	return funcao(numero1, numero2)
}
