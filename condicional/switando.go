package condicional

import (
	"fmt"
	"runtime"
	"time"
)

//TestaSwitch ae
func TestaSwitch() {

	num := 1

	fmt.Println("O numero ", num, "escreve assim: ")

	switch num {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")

	}

	fmt.Println("Voce e da familia do Unix")
	switch runtime.GOOS {
	case "dwarwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim")
	default:
		fmt.Println("Deixa pra la")

	}

	fmt.Println("O que fazer hoje? ")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Vai dormir")
	default:
		fmt.Println("Vamo trabalhar")

	}

	num = 10

	fmt.Println("Esse numero cabe num digito? ")

	switch {
	case num < 10:
		fmt.Println("Sim")
	case num >= 10 && num < 100:
		fmt.Println("Cabee em dois...")
	case num >= 100:
		fmt.Println("Grande demais deixa proa la")
	}
}
