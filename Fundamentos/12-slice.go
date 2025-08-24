package main

import (
	"fmt"
)

func main() {
	var numeros []int
	//add elementos slice
	numeros = append(numeros, 10)
	numeros = append(numeros, 20)
	numeros = append(numeros, 30, 40, 50)
	fmt.Println("Slice de números:", numeros)

	frutas := []string{"maçã", "banana", "laranja", "uva"}
	fmt.Println("Slice de frutas:", frutas)

	subslice := frutas[1:3] // Pega elementos do índice 1 ao 2 (3 não incluso)
	fmt.Println("Subslice de frutas:", subslice)

	frutas[0] = "abacaxi"
	fmt.Println("Slice de frutas após modificação:", frutas)
	fmt.Println("Subslice de frutas após modificação:", subslice)
}
