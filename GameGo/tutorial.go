package main

import (
	"fmt"
	"strings"
)

func main() {
	/*fmt.Println("welcome to my quiz game")

	var name string = "Kauã"
	fmt.Println(name)

	var age int = 22
	fmt.Println(age)

	var height float64 = 1.82
	fmt.Println(height)

	var name1 uint = 123
	fmt.Println(name1)

	name2 := "Pedr0"
	fmt.Println(name2)
	fmt.Printf("the type of name2 is %T\n", name2)
	fmt.Printf("Hello %v\n", name2)

	var nome3 string
	fmt.Println("Digite seu nome: ")
	// Passa o endereço da variável (&) para o Scanln preencher
	fmt.Scanln(&nome3)
	fmt.Println("Olá,", nome3)*/

	fmt.Println("🎮 Welcome to my quiz game! 🎮")
	fmt.Println(strings.Repeat("=", 40))

	// Nome do jogador
	var name string
	fmt.Print("Your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n\n", name)

	// Idade
	var age int
	fmt.Print("Your age: ")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("❌ You are not old enough to play this game.")
		return
	} else {
		fmt.Println("✅ Your age is", age, "- you can play! Good luck!")
	}

	score := 0
	totalQuestions := 3

	// Pergunta 1
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Question 1: What is better for you?")
	fmt.Println("1 - RTX 3080")
	fmt.Println("2 - RTX 3090")

	var answer1 string
	fmt.Scanln(&answer1)
	answer1 = strings.ToUpper(strings.TrimSpace(answer1))

	if answer1 == "RTX 3090" || answer1 == "2" {
		fmt.Println("✅ Correct!")
		score++
	} else {
		fmt.Println("❌ Incorrect! The answer is RTX 3090")
	}

	// Pergunta 2
	fmt.Println(strings.Repeat("-", 40))
	fmt.Print("Question 2: How many cores does the Ryzen 9 3950X have? ")
	var cores int
	fmt.Scanln(&cores)

	if cores == 16 {
		fmt.Println("✅ Correct!")
		score++
	} else {
		fmt.Printf("❌ Incorrect! The Ryzen 9 3950X has 16 cores.\n")
	}

	// Pergunta 3 (adicional)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Print("Question 3: What company makes the RTX graphics cards? ")
	var company string
	fmt.Scanln(&company)
	company = strings.ToLower(strings.TrimSpace(company))

	if company == "nvidia" {
		fmt.Println("✅ Correct!")
		score++
	} else {
		fmt.Println("❌ Incorrect! The answer is NVIDIA")
	}

	// Resultado final
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("\n📊 Final Score: %d/%d\n", score, totalQuestions)

	percentage := float64(score) / float64(totalQuestions) * 100
	fmt.Printf("📈 Percentage: %.1f%%\n", percentage)

	if percentage >= 70 {
		fmt.Println("🏆 Congratulations! You passed!")
	} else {
		fmt.Println("📚 Keep studying and try again!")
	}
}
