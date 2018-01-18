package estrutura

//Carro teste com funcoes
type Carro struct {
	Marca  string `json:"marca"`
	Modelo string `json:"Modelo"`
	valor  float64
}

//SetValor seta como se fosse encapsulamento
func (c *Carro) SetValor(valor float64) {
	c.valor = valor
}

//GetValor obtem como se fosse encapuslamento de valors privados
func (c *Carro) GetValor() float64 {
	return c.valor
}
