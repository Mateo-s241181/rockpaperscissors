package player

import (
	"fmt"
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
	// solution:begin
	return p.Name
	// solution:end
}

// GetScore liefert den Punktestand des Spielers.
func (p Human) GetScore() int {
	// solution:begin
	return p.Score
	// solution:end
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Human) IncrementScore() {
	// solution:begin
	p.Score++
	// solution:end
}

// GetMove liefert einen Zug des Spielers.
func (p Human) GetMove() values.Value {
	// solution:begin
	fmt.Printf("%s, was wählst du?\n", p.GetName())
	fmt.Println("  1: Stein")
	fmt.Println("  2: Papier")
	fmt.Println("  3: Schere")
	var v int
	fmt.Scanln(&v)
	return values.Value((v - 1) % 3)
	// solution:end
}
