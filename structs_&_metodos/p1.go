package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	idade int
	email string
	sexo  string
}

func main() {
	p1 := Pessoa{
		Nome:  "Kauã",
		idade: 22,
		email: "kaua@example.com",
		sexo:  "Masculino",
	}

	fmt.Println(p1)
}
