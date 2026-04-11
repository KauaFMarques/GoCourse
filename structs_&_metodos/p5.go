package main

import (
	"fmt"
)

type Guerreiro struct {
	Nome string
	Vida int
}

func (g Guerreiro) Status() {
	fmt.Printf("O guerreiro %s está com %d de vida\n", g.Nome, g.Vida)
}

// Adicione o * para ser um ponteiro e REALMENTE tirar vida do original
func (g *Guerreiro) TomarDano(dano int) {
	g.Vida -= dano
	if g.Vida < 0 {
		g.Vida = 0
	}
	fmt.Printf("%s tomou %d de dano!\n", g.Nome, dano)
}

func main() {
	player := Guerreiro{Nome: "Aragorn", Vida: 100}

	player.Status()
	player.TomarDano(30)
	player.Status() // Agora vai aparecer 70!
}
