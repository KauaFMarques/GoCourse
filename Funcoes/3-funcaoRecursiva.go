package main

import (
	"fmt"
)

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func totalSomatorio(n int) int {
	if n == 0 {
		return 0
	}
	return n + totalSomatorio(n-1)
}

func main() {
	var num int
	fmt.Print("Digite um número para calcular o fatorial: ")
	fmt.Scan(&num)
	resultado := fatorial(num)
	fmt.Printf("O fatorial de %d é %d\n", num, resultado)
	somatorio := totalSomatorio(num)
	fmt.Printf("O somatório de 0 até %d é %d\n", num, somatorio)
}
