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

	fmt.Println(strings.ToUpper(moviedescription))
	fmt.Println(strings.ToLower(moviedescription))
	fmt.Println(strings.Title(moviedescription))
	fmt.Println(strings.Index(moviedescription, "Kosinski"))
}
