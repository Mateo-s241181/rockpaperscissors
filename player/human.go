package player

import (
	"fmt"
	"rockpaperscissors/values"
	"strings"
)

type Human struct {
	Name  string
	Score int
}

// NewHuman erzeugt einen neuen Spieler mit dem Namen n.
func NewHuman(n string) *Human {
	return &Human{Name: n, Score: 0}
}

// GetName liefert den Namen des Spielers als String.
func (p Human) GetName() string {
	return p.Name
}

// GetScore liefert den Punktestand des Spielers.
func (p Human) GetScore() int {
	return p.Score
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Human) IncrementScore() {
	p.Score++
}

// GetMove liefert einen Zug des Spielers.
func (p Human) GetMove() values.Value {

	//String mit Int in Verbindung setzen
	movesMap := map[string]int{
		"schere": 0,
		"stein":  1,
		"papier": 2,
	}

	var move string

	//Spieler nach Input fragen
	fmt.Printf("%v, du bist dran! \nSchere, Stein, Papier!\n", p.GetName())
	fmt.Scan(&move)

	//den string in Lowercase konvertieren
	move = strings.ToLower(move)

	//bei falschem Input Fehlernachricht ausgeben und Prozess wiederholen
	if move != "schere" && move != "stein" && move != "papier" {
		fmt.Println("Du musst zwischen Schere, Stein und Papier wählen.")
		return p.GetMove()
	}

	//Valuewert returnen
	return values.Value(movesMap[move])
}
