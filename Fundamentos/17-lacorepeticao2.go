package main

import (
	"fmt"
)

func main() {
	notas := []float64{6.5, 7.8, 9.1, 4.3, 5.6}
	fmt.Println("calculo da media com slice:")

	var total float64
	for _, nota := range notas {
		total += nota
		fmt.Println("nota:", nota)
	}
	fmt.Println("Soma das notas:", total)
	media := total / float64(len(notas))
	fmt.Printf("Média: %.2f\n", media)
}
