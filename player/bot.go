package player

import (
	"math/rand"
	"rockpaperscissors/values"
)

type Bot struct {
	Name  string
	Score int
}

// NewBot erzeugt einen neuen Spieler mit dem Namen n.
func NewBot(n string) *Bot {
	return &Bot{Name: n, Score: 0}
}

// GetName liefert den Namen des Spielers als String.
func (p Bot) GetName() string {
	return p.Name
}

// GetScore liefert den Punktestand des Spielers.
func (p Bot) GetScore() int {
	return p.Score
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Bot) IncrementScore() {
	p.Score++
}

// GetMove liefert einen Zug des Spielers.
func (p Bot) GetMove() values.Value {

	return values.Value(rand.Intn(3))
}
