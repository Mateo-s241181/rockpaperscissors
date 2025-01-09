package player

import "rockpaperscissors/values"

type Player interface {
	// GetName liefert den Namen des Spielers als String.
	GetName() string
	// GetScore liefert den Punktestand des Spielers.
	GetScore() int
	// IncrementScore erhöht den Punktestand des Spielers um eins.
	IncrementScore()
	// GetMove liefert einen Zug des Spielers.
	GetMove() values.Value
}
