package player

import (
	"rockpaperscissors/values"
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
	// TODO
	return ""
}

// GetScore liefert den Punktestand des Spielers.
func (p Human) GetScore() int {
	// TODO
	return 0
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Human) IncrementScore() {
	// TODO
}

// GetMove liefert einen Zug des Spielers.
func (p Human) GetMove() values.Value {
	// TODO
	return values.Rock
}
