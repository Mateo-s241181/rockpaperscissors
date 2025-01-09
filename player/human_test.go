package player

import "fmt"

func ExampleHuman_IncrementScore() {
	p := NewHuman("Alice")
	p.IncrementScore()
	p.IncrementScore()
	fmt.Println(p.GetScore())

	// Output:
	// 2
}

func ExampleHuman_GetName() {
	p := NewHuman("Alice")

	fmt.Println(p.GetName())

	// Output:
	// Alice
}
