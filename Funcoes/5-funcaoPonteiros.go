package main

import (
	"fmt"
)

func alterValue(x *int) {
	*x = *x * 2
}

func main() {
	num := 10
	fmt.Println("Valor inicial:", num)
	alterValue(&num)
	fmt.Println("Valor alterado:", num)
}
