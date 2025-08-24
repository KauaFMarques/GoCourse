package main

import (
	"fmt"
)

func main() {
	var nota float32
	fmt.Print("Digite sua nota: ")
	fmt.Scan(&nota)

	switch {
	case nota < 0 || nota > 10:
		fmt.Println("Nota inválida")
	case nota >= 7:
		fmt.Println("Aprovado")
	case nota >= 5:
		fmt.Println("Recuperação")
	default:
		fmt.Println("Reprovado")
	}
}
