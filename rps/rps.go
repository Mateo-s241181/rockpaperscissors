package main

import (
	"fmt"
	"rockpaperscissors/game"
)

func main() {
	var n1 string
	var n2 string
	fmt.Print("Bitte den Namen von Spieler 1 eingeben: ")
	fmt.Scanln(&n1)
	fmt.Print("Bitte den Namen von Spieler 2 eingeben: ")
	fmt.Scanln(&n2)

	g := game.NewBotGame(n1, n2)

	p2_human := "N"
	fmt.Printf("Wird %s von einem Menschen gespielt? Falls ja, antworte mit \"Ja\" oder \"J\"", n2)
	fmt.Scanln(&p2_human)
	if p2_human == "J" || p2_human == "Ja" {
		fmt.Println("Zweispielerspiel")
		g = game.NewHumanGame(n1, n2)
	} else {
		fmt.Println("Einzelspielerspiel")
	}

	for i := 0; i < 4; i++ {
		g.Play(i + 1)
	}

	fmt.Printf("\n\n------------------------------------------\n")

	if g.Player1.GetScore() > g.Player2.GetScore() {
		fmt.Printf("%v hat das Spiel gewonnen!", g.Player1.GetName())
	}
	if g.Player2.GetScore() > g.Player1.GetScore() {
		fmt.Printf("%v hat das Spiel gewonnen!", g.Player2.GetName())
	}
	if g.Player2.GetScore() == g.Player1.GetScore() {
		fmt.Printf("Das Spielergebnis ist ein Unentschieden...")
	}

	fmt.Printf("\n------------------------------------------")
}
