// aprendendo a fazer um contador
package main

import (
	"fmt"
)

type Contador struct {
	Valor int
}

func (c *Contador) Incrementar() {
	c.Valor++
}

func (c *Contador) Decrementar() {
	c.Valor--
}

func (c *Contador) Zerar() {
	c.Valor = 0
}

func (c *Contador) Exibir() {
	fmt.Printf("O valor do contador é: %d\n", c.Valor)
}

func main() {
	contador := Contador{Valor: 0}

	contador.Exibir()
	contador.Incrementar()
	contador.Exibir()
	contador.Incrementar()
	contador.Exibir()
	contador.Decrementar()
	contador.Exibir()
	contador.Zerar()
	contador.Exibir()

}
