package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(10) + 1 // Número aleatório entre 1 e 10
	fmt.Println("Tente adivinhar o número entre 1 e 10:")
	var guess int
	for {
		fmt.Print("Digite seu palpite ou 0 para sair: ")
		fmt.Scan(&guess)
		if guess == 0 {
			fmt.Println("Jogo encerrado. O número era:", target)
			break
		} else if guess < target {
			fmt.Println("Tente um numero maior. Vc digitou um muito baixo")
		} else if guess > target {
			fmt.Println("Tente um numero menor. Vc digitou um muito alto")
		} else {
			fmt.Println("Parabéns! Você acertou o número:", target)
			break
		}
	}
}
