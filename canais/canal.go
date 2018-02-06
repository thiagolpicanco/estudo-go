package canais

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

//TestaCanais testaCanais, um canal serve para mandar um sinal(msg) para outra rotina
func TestaCanais() {
	go pinger(irc)
	go ponger(irc)
	go epa(sms)
	go opa(sms)
	go imprimir()
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

func epa(canal chan string) {
	for {
		canal <- "Epa"
	}
}

func opa(canal chan string) {
	for {
		canal <- "Opa"
	}
}

func imprimir() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		case msg = <-sms:
			fmt.Println("chegou sms: ", msg)
		}
		time.Sleep(time.Second * 1)
	}

}
