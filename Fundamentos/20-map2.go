package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "O rato roeu a roupa do rei de Roma Roma"
	words := strings.Fields(text)
	fmt.Println("Palavras na string:", words)
	fmt.Println("Texto original:", text)

	wordCount := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	//imprimir o mapa de contagem de palavras
	fmt.Println("Contagem de palavras:")
	for word, count := range wordCount {
		fmt.Printf("word: %s: contagem: %d\n", word, count)
	}
}
