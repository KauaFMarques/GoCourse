package main

import "fmt"

func main() {
	var nome string
	var idade int
	var altura float32
	var maiorIdade bool

	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	fmt.Print("Digite sua altura (em metros): ")
	fmt.Scan(&altura)

	maiorIdade = idade >= 18
	fmt.Printf("\nOlá, %s!\n", nome)
	fmt.Printf("Idade: %d anos\n", idade)
	fmt.Printf("Altura: %.2f metros\n", altura)
	fmt.Printf("Maior de idade: %t\n", maiorIdade)
}
