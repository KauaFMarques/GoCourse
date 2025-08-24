package main

import (
	"fmt"
)

type Cliente struct {
	Nome   string
	Idade  int
	Altura float32
	Peso   float32
	Casado bool
	Cidade string
}

func atualizarCliente(c *Cliente, novoNome string, novaIdade int, novaAltura float32, novoPeso float32, novoEstadoCivil bool, novaCidade string) {
	c.Nome = novoNome
	c.Idade = novaIdade
	c.Altura = novaAltura
	c.Peso = novoPeso
	c.Casado = novoEstadoCivil
	c.Cidade = novaCidade

}

func main() {
	cliente := Cliente{
		Nome:   "Ana",
		Idade:  28,
		Altura: 1.65,
		Peso:   60.0,
		Casado: false,
		Cidade: "Rio de Janeiro",
	}
	fmt.Println("Antes da atualização:")
	fmt.Printf("%+v\n", cliente)
	atualizarCliente(&cliente, "Ana Silva", 29, 1.66, 62.0, true, "São Paulo")
	fmt.Println("Depois da atualização:")
	fmt.Printf("%+v\n", cliente)

}
