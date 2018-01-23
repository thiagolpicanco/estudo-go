package array

import (
	"fmt"
)

//TestaArrays testa array
func TestaArrays() {

	var teste [3]int
	teste[0] = 1
	teste[1] = 2
	teste[2] = 3

	for indice, valor := range teste {
		fmt.Println(indice, valor)
	}

	fmt.Println("Capacidade", len(teste))

	array2 := [2]string{"James", "Carlos"}

	fmt.Println(array2)

	array3 := [...]string{"Lisboa", "Rio", "São Paulo"}
	for indice, valor := range array3 {
		fmt.Printf("Indice %v\n , valor: %s\n", indice, valor)
	}
}
