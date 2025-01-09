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
	// Hinweis: Greifen Sie mit p.Name auf den Namen des Spielers zu.

	// TODO
	return ""
}

// GetScore liefert den Punktestand des Spielers.
func (p Human) GetScore() int {
	// Hinweis: Greifen Sie mit p.Score auf den Punktestand des Spielers zu.

	// TODO
	return 0
}

// IncrementScore erhöht den Punktestand des Spielers um eins.
func (p *Human) IncrementScore() {
	// Hinweis: Greifen Sie mit p.Score auf den Punktestand des Spielers zu.
	//          Sie können p.Score wie eine reguläre Variable verwenden.

	// TODO
}

// GetMove liefert einen Zug des Spielers.
func (p Human) GetMove() values.Value {
	// Hinweis: Fordern Sie den Spieler mit fmt.Printf() auf, einen Zug zu wählen.
	//          Geben Sie dem Spieler die Optionen "Stein", "Papier" und "Schere".
	//          Die Optionen können z.B. mit Zahlen 1, 2 und 3 verknüpft werden.
	//          Verwenden Sie fmt.Scanln() um die Eingabe des Spielers zu lesen.

	// TODO
	return values.Rock
}
