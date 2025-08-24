package main

import (
	"fmt"
)

type Carro struct {
	Marca  string
	Ano    int
	Modelo string
	Cor    string
}

func main() {
	meuCarro := Carro{
		Marca:  "Toyota",
		Ano:    2020,
		Modelo: "Corolla",
		Cor:    "Prata",
	}

	fmt.Println("Detalhes do meu carro:")
	fmt.Printf("Marca: %s\n", meuCarro.Marca)
	fmt.Printf("Ano: %d\n", meuCarro.Ano)
	fmt.Printf("Modelo: %s\n", meuCarro.Modelo)
	fmt.Printf("Cor: %s\n", meuCarro.Cor)
}
