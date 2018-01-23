package slice

import (
	"fmt"
)

//TestaSlice ae
func TestaSlice() {

	var nums []int

	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 3)

	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}

	fmt.Println(capitais, len(capitais), cap(capitais))
	//ADICIONAR MAIS UM ITEM
	capitais = append(capitais, "Brasilia")

	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 4)
	cidades[0] = "Saquarema"
	cidades[1] = "Toquio"
	cidades[2] = "Arraial"
	cidades[3] = "Buzios"

	fmt.Println(cidades, len(cidades), cap(cidades))
	fmt.Println(cidades)

	//FAZENDO O CORTE
	//primeiro e segundo
	//O PRIMEIRO ITEM COMECA COM INDICE 0
	// O SEGUNDA COMECA COM INDICE 1
	teste := cidades[2:4]
	fmt.Println(teste)

	//INICIA DO ZERO
	teste2 := cidades[:4]

	fmt.Println(teste2)

	teste3 := cidades[len(cidades)-2:]
	fmt.Println(teste3)

	//REMOVER ITEM DO SLICE

	indiceRemover := 2
	temp := cidades[:indiceRemover]
	temp = append(temp, cidades[indiceRemover+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)
}
