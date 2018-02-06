package canais

import (
	"fmt"
	"time"
)

//TestaCanais testaCanais, um canal serve para mandar um sinal(msg) para outra rotina
func TestaCanais() {
	var canal chan string
	canal = make(chan string)
	go pinger(canal)
	go ponger(canal)
	go imprimir(canal)
	var entrada string
	fmt.Scanln(&entrada)

}

func pinger(canal chan string) {
	for {
		canal <- "ping"
	}

}
func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func imprimir(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

}
