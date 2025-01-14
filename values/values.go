package values

type Value int

const (
	// iota weist allen Values einen Integer Wert zu, der bei jeder Zuweisung inkrementiert wird
	Schere Value = iota //0
	Stein               //1
	Papier              //2
)

// String liefert den Wert als String.
func (v Value) String() string {

	s := ""

	switch v {
	case 0:
		s = "Schere"
	case 1:
		s = "Stein"
	case 2:
		s = "Papier"
	}

	return s
}

// Beats gibt an, ob der Wert v den Wert w schlägt.
func (v Value) Beats(w Value) bool {

	if v == w {
		return false
	}

	Beats := false

	switch v {

	//rock only beats scissors
	case 0:

		if w == 2 {
			Beats = true
		}

	//paper only beats rock
	case 1:

		if w == 0 {
			Beats = true
		}

	//scissors only beat paper
	case 2:

		if w == 1 {
			Beats = true
		}

	}
	return Beats
}
