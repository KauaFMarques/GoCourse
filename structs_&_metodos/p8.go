package main

import (
	"fmt"
)

type Produto struct {
	Nome  string
	Preco float64
}

type Carrinho struct {
	Produtos []Produto
}

// Método Total: percorre o slice e soma os preços
func (c Carrinho) Total() float64 {
	soma := 0.0
	// O 'range' percorre o slice de produtos
	for _, p := range c.Produtos {
		soma += p.Preco
	}
	return soma
}

func main() {
	p1 := Produto{Nome: "Teclado", Preco: 150.50}
	p2 := Produto{Nome: "Mouse", Preco: 80.00}

	meuCarrinho := Carrinho{
		Produtos: []Produto{p1, p2},
	}

	fmt.Printf("Total do carrinho: R$ %.2f\n", meuCarrinho.Total())
}
