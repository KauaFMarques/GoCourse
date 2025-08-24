package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)
	if idade < 18 {
		fmt.Println("Você é menor de idade.")
	} else {
		fmt.Println("Você é maior de idade.")
	}

	var nota float32
	fmt.Println("digite sua nota: ")
	fmt.Scan(&nota)
	if nota > 10 || nota < 0 {
		fmt.Println("Nota invalida")
	} else if nota >= 7 {
		fmt.Println("Aprovado")
	} else if nota >= 5 {
		fmt.Println("Recuperação")
	} else {
		fmt.Println("Reprovado")
	}
}
