// conversor de temperatura
package main

import (
	"fmt"
)

type Celsius struct {
	temperatura float64
}

func (c *Celsius) ParaFahrennheit() float64 {
	return (c.temperatura * 9 / 5) + 32
}

func main() {
	minhaTemperatura := Celsius{temperatura: 25.0}
	f := minhaTemperatura.ParaFahrennheit()

	fmt.Printf("%.2f°C equivalem a %.2f°F\n", minhaTemperatura.temperatura, f)
}
