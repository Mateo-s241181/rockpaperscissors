package game

import (
	"fmt"
	"rockpaperscissors/player"
)

type Game struct {
	Player1 player.Player
	Player2 player.Player
}

// NewHumanGame erzeugt ein neues Spiel mit zwei menschlichen Spielern mit den Namen n1 und n2.
func NewHumanGame(n1, n2 string) *Game {
	return &Game{
		Player1: player.NewHuman(n1),
		Player2: player.NewHuman(n2),
	}
}

// NewBotGame erzeugt ein neues Spiel mit einem menschlichen und einem Bot-Spieler
// mit den Namen n1 und n2.
func NewBotGame(n1, n2 string) *Game {
	return &Game{
		Player1: player.NewHuman(n1),
		Player2: player.NewBot(n2),
	}
}

// Play spielt ein Spiel mit den beiden Spielern.
// D.h. die Funktion fragt jeden Spieler nach einem Zug und ermittelt den Gewinner.
// Erhöht den Punktestand des Gewinners und zeigt die Punkte an.
func (g *Game) Play(counter int) {

	fmt.Println("---------------")
	fmt.Printf("    Runde%d    \n", counter)
	fmt.Println("---------------")

	counter++

	v1 := g.Player1.GetMove()
	v2 := g.Player2.GetMove()

	fmt.Printf("%v hat %v gewählt.\n", g.Player1.GetName(), v1)
	fmt.Printf("%v hat %v gewählt.\n", g.Player2.GetName(), v2)

	//Falls beide gleich sind soll kein score inkrementiert werden
	if v1 != v2 {

		//Je nach dem wer gewinnt wird der score des gewinners inkrementiert
		switch v1.Beats(v2) {
		case true:
			g.Player1.IncrementScore()
		case false:
			g.Player2.IncrementScore()
		}
	}

	g.PrintScores()
}

// PrintScores gibt den Punktestand beider Spieler auf die Konsole aus.
func (g Game) PrintScores() {
	fmt.Println("Punktestand:")
	fmt.Printf("  %v: %d\n", g.Player1.GetName(), g.Player1.GetScore())
	fmt.Printf("  %v: %d\n\n", g.Player2.GetName(), g.Player2.GetScore())
}
