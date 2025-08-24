package main

import (
	"fmt"
	"strings"
)

func main() {
	movie := "top gun: maverick"
	movie2 := "Top gun: Maverick"

	fmt.Println(movie == movie2)

	moviedescription := `é um filme de ação e drama americano de 2022 dirigido por Joseph Kosinski.`

	fmt.Println(moviedescription)

	line := "O filme é uma sequência do filme de 1986 Top Gun."
	fmt.Println(strings.Repeat(line+"\n", 3))

	fmt.Println(strings.Contains(movie, "gun"))
}
