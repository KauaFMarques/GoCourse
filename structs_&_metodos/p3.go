// embedded
package main

import (
	"fmt"
)

// 1. Defina a struct base fora do main (ou antes de usar)
type Pessoa struct {
	Nome  string
	Idade int
}

// 2. Defina a struct que faz o embedding
type Funcionario struct {
	Pessoa  // Campo embutido (promovido)
	Salario float64
}

func main() {
	// 3. Mude o nome da variável (não use 'func')
	funcionario := Funcionario{
		Pessoa:  Pessoa{Nome: "Pedro", Idade: 28},
		Salario: 5000.00,
	}

	fmt.Println(funcionario.Nome)  // Acesso direto via promoção de campos
	fmt.Println(funcionario.Idade) // 28
}
