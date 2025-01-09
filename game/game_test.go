package game

func ExampleGame_PrintScores() {
	g := NewHumanGame("A", "B")
	g.PrintScores()

	g.Player1.IncrementScore()
	g.PrintScores()

	g.Player2.IncrementScore()
	g.PrintScores()

	// Output:
	// Punktestand:
	//   A: 0
	//   B: 0
	//
	// Punktestand:
	//   A: 1
	//   B: 0
	//
	// Punktestand:
	//   A: 1
	//   B: 1
}
