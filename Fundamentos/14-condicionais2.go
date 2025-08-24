package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Println("Informe um numero:")
	fmt.Scan(&numero)
	if numero > 0 {
		fmt.Println("O número é positivo")
	} else if numero < 0 {
		fmt.Println("O número é negativo")
	} else if numero%2 == 0 {
		fmt.Println("O número é par")
	} else if numero%2 != 0 {
		fmt.Println("O número é impar")
	} else {
		fmt.Println("O número é zero")
	}
}
