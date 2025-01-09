package player

import "fmt"

func ExampleBot_IncrementScore() {
	p := NewBot("Alice")
	p.IncrementScore()
	p.IncrementScore()
	fmt.Println(p.GetScore())

	// Output:
	// 2
}

func ExampleBot_GetName() {
	p := NewBot("Alice")

	fmt.Println(p.GetName())

	// Output:
	// Alice
}
