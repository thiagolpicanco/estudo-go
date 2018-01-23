package estrutura

import (
	"errors"
)

//Carro teste com funcoes
type Carro struct {
	Marca  string `json:"marca"`
	Modelo string `json:"Modelo"`
	Valor  float64
}

//ErrValorCarroInvalido joga erro
var ErrValorCarroInvalido = errors.New("O valor informado não é valido para um caro")

//ErrValorCarroAlto joga erro
var ErrValorCarroAlto = errors.New("O valor informado muito alto")

//SetValor seta como se fosse encapsulamento
func (c *Carro) SetValor(valor float64) (err error) {
	err = nil

	if valor < 1000 {
		err = ErrValorCarroInvalido
		return
	} else if valor > 100000 {
		err = ErrValorCarroAlto
	}
	c.Valor = valor
	return
}

//GetValor obtem como se fosse encapuslamento de valors privados
func (c *Carro) GetValor() float64 {
	return c.Valor
}
