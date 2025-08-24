package main

import (
	"fmt"
)

func main() {
	var notas = [3]float32{8.5, 9.0, 7.5}
	fmt.Println("Notas dos alunos:", notas)

	var soma = notas[0] + notas[1] + notas[2]
	media := soma / float32(len(notas))
	fmt.Printf("Média da turma: %.2f\n", media)
}
