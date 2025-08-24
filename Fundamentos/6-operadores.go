package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	//soma
	soma := num1 + num2
	fmt.Println("Soma:", soma)
	//subtração
	subtracao := num1 - num2
	fmt.Println("Subtração:", subtracao)
	//multiplicação
	multiplicacao := num1 * num2
	fmt.Println("Multiplicação:", multiplicacao)
	//divisão
	divisao := num1 / num2
	fmt.Println("Divisão:", divisao)

	//atribuição
	num3 := 10
	num3 += 5 // num3 = num3 + 5
	fmt.Println("Atribuição (num3 += 5):", num3)

	//comparação
	igual := num1 == num2
	fmt.Println("Igualdade (num1 == num2):", igual)
	diferente := num1 != num2
	fmt.Println("Diferença (num1 != num2):", diferente)
}
