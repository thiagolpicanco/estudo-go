package loops

import (
	"fmt"
)

//TestaLoop testa loop
func TestaLoop() {

	//TRADICIONAL
	for i := 0; i < 10; i++ {
		fmt.Println("QUal o valor de i", i)
	}

	//FOR ESTILO WHILE
	valor := 0
	teste := true
	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O numero é ", valor)
	}

	//LOOP ATE O BREAK
	for {
		valor--
		fmt.Println("O numero é ", valor)
		if valor == 0 {
			break
		}
	}

	texto := "Vasco da Gama"

	//ESTILO FOR EACH
	for indice, letra := range texto {
		fmt.Printf("Texto [%d] = %q\r\n", indice, letra)
	}

}
