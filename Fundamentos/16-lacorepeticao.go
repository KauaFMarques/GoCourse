package main

import (
	"fmt"
)

func main() {
	nomes := []string{"Ana", "Bruno", "Carla", "Daniel"}
	for i, nome := range nomes {
		fmt.Printf("Índice: %d, Nome: %s\n", i, nome)
		if len(nome) > 4 {
			fmt.Printf("O nome %s tem mais de 4 letras.\n", nome)
		}
	}
}
