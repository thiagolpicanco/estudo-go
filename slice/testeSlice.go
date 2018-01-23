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

	cidades = append(cidades, "Cabo Frio")

	fmt.Println(cidades, len(cidades), cap(cidades))
	fmt.Println(cidades)
}
