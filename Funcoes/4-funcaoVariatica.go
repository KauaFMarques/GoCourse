package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	sumTotal := 0
	for _, n := range numbers {
		sumTotal += n
	}
	return sumTotal
}

//2-função para apresentação de cursos variáticos

func presentation(data map[string]string) {
	for key, value := range data {
		fmt.Printf("%s - %s\n", key, value)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum())
	presentation(map[string]string{
		"curso1": "Go",
		"curso2": "JavaScript",
		"curso3": "Python",
	})
}
