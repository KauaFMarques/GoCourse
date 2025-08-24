package main

import "fmt"

func main() {
	var idade int = 30
	var altura float32 = 1.75
	var nome string = "João"
	var maiorIdade bool = idade >= 18
	fmt.Println("Dados Pessoais: ")
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Nome:", nome)
	fmt.Println("Maior de Idade:", maiorIdade)
	fmt.Printf("Tipo da variável idade: %T\n", idade)
}
