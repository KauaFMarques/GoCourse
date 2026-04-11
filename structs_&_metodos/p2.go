// structs aninhadas
package main

import (
	"fmt"
)

type Endereco struct {
	Rua    string
	Cidade string
	CEP    string
}

type Cliente struct {
	Nome     string
	Endereco Endereco // composição
}

func main() {
	cliente := Cliente{
		Nome: "Ana",
		Endereco: Endereco{
			Rua:    "Rua A",
			Cidade: "São Paulo",
			CEP:    "12345-678",
		},
	}

	fmt.Println(cliente.Endereco.Cidade) // São Paulo
}
