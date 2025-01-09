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
	// solution:begin
	return p.Name
	// solution:end
}

// GetScore liefert den Punktestand des Spielers.
func (p Bot) GetScore() int {
	// solution:begin
	return p.Score
	// solution:end
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Bot) IncrementScore() {
	// solution:begin
	p.Score++
	// solution:end
}

// GetMove liefert einen Zug des Spielers.
func (p Bot) GetMove() values.Value {
	// solution:begin
	return values.Value(rand.Intn(3))
	// solution:end
}
