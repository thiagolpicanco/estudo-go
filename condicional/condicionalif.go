package condicional

import (
	"fmt"
)

//TestaIF faz isso
func TestaIF() {

	meses := 1
	devendo := true
	cidade := "Saqua"

	if devendo {
		fmt.Println("devedor")
	}

	if !devendo {
		fmt.Println("adimnplente")
	}

	if cidade == "Saqua" {
		fmt.Println("Mora legal")
	}

	if meses > 1 {
		fmt.Println("mes maior q 1")
	} else if meses == 1 {
		fmt.Println("Mesu igual a 1")
	} else {
		fmt.Println("mes menor q 1")

	}

	if descricao, status := clienteEstaDevendo(meses); status {
		fmt.Println(descricao)
		return
	}

	//LEMBRANDO QUE AS VARIAVEIS DECLARADAS DENTRO DO IF SO FUNCIONAM LA
}

func clienteEstaDevendo(meses int) (descricao string, status bool) {

	if meses > 0 {
		status = true
		descricao = "Devendo"
		return
	}
	descricao = "Cliente esta adimplente"
	return
}
