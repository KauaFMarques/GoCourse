package main

import "fmt"

func fullName(firstName string, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func someNumbers(a int, b int) int {
	return a + b
}

func adress(country string) {
	if country == "" {
		country = "Brasil"
	}
	fmt.Println("País:", country)
}

func main() {
	fmt.Println(fullName("João", "Silva"))
	fmt.Println(someNumbers(10, 20))
	adress("")
}
