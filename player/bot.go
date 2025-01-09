package player

import (
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
	// Hinweis: Greifen Sie mit p.Name auf den Namen des Spielers zu.

	// TODO
	return ""
}

// GetScore liefert den Punktestand des Spielers.
func (p Bot) GetScore() int {
	// Hinweis: Greifen Sie mit p.Score auf den Punktestand des Spielers zu.

	// TODO
	return 0
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Bot) IncrementScore() {
	// Hinweis: Greifen Sie mit p.Score auf den Punktestand des Spielers zu.
	//          Sie können p.Score wie eine reguläre Variable verwenden.

	// TODO
}

// GetMove liefert einen Zug des Spielers.
func (p Bot) GetMove() values.Value {
	// Hinweis: Verwenden Sie die Funktion rand.Intn(3) um eine Zufallszahl zwischen 0 und 2 zu erzeugen.
	//          Verwenden Sie dann values.Value() um die Zufallszahl in einen Wert umzuwandeln.

	// TODO
	return values.Rock
}
