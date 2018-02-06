package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var orquestrador sync.WaitGroup

func TestaRotinas() {

	arrayRegiaoLagos := [...]string{"Arraial", "Saquarema", "Cabo Frio", "Buzios"}

	arrayMetropolitan := [...]string{"Niteroi", "Caxias", "Rio de Janeiro", "Sao Goncalo"}
	orquestrador.Add(2)
	go imprimir(arrayRegiaoLagos)
	go imprimir(arrayMetropolitan)
	orquestrador.Wait()

}

func imprimir(array [4]string) {
	fmt.Println("Comecou as ", time.Now())
	for _, valor := range array {
		fmt.Printf("cidade: %s\n", valor)
	}
	orquestrador.Done()

}
