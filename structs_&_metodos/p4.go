// aprendendo metodos
package main

import (
	"fmt"
)

type Retangulo struct {
	Largura float64
	Altura  float64
}

// Método para calcular a area sem alterar dados da struct retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Método com receiver por ponteiro (modifica o original)
func (r *Retangulo) Redimensionar(fator float64) {
	r.Largura *= fator
	r.Altura *= fator
}

func main() {
	ret := Retangulo{Largura: 10, Altura: 5}

	fmt.Printf("Área inicial: %.2f\n", ret.Area())

	// Chamando o método que usa ponteiro para dobrar o tamanho
	ret.Redimensionar(2)

	fmt.Printf("Nova largura: %.2f\n", ret.Largura)
	fmt.Printf("Nova área: %.2f\n", ret.Area())
}
