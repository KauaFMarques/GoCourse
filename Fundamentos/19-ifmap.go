package main

import (
	"fmt"
)

func main() {
	estudantes := map[string]float64{
		"Alice": 8.5,
		"Bob":   6.0,
		"Carol": 9.2,
		"Dave":  4.8,
	}
	//fmt.Println("Mapa de estudantes e suas notas: \n", estudantes)
	fmt.Println("Classificação dos estudantes")

	for nome, nota := range estudantes {
		if nota < 0 || nota > 10 {
			fmt.Printf("Nota inválida para %s: %.2f\n", nome, nota)
		} else if nota >= 7 {
			fmt.Printf("%s está Aprovado com nota %.2f\n", nome, nota)
		} else if nota >= 5 {
			fmt.Printf("%s está em Recuperação com nota %.2f\n", nome, nota)
		} else {
			fmt.Printf("%s está Reprovado com nota %.2f\n", nome, nota)
		}
	}
}
