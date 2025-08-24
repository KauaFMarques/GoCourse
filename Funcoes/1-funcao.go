package main

import (
	"fmt"
)

func welcome() {
	fmt.Println("Bem-vindo ao sistema de filmes!")
}

func createMovie() {
	var title string
	var year int
	fmt.Println("Digite o título do filme:")
	fmt.Scan(&title)
	fmt.Println("Digite o ano de lançamento do filme:")
	fmt.Scan(&year)
	fmt.Printf("Filme criado: %s (%d)\n", title, year)
}

func clearBuffer() {
	var discard string
	fmt.Scanln(&discard)
}

func main() {
	welcome()
	fmt.Println("Olá, Mundo!")
	createMovie()
	clearBuffer()
	fmt.Println("Programa encerrado.")
}
