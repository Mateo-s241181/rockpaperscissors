package values

import "fmt"

func ExampleValue_String() {
	schere := Schere
	stein := Stein
	papier := Papier

	fmt.Println(schere)
	fmt.Println(stein)
	fmt.Println(papier)

	// Output:
	// Schere
	// Stein
	// Papier
}

func ExampleValue_Beats() {
	fmt.Println(Schere.Beats(Schere))
	fmt.Println(Schere.Beats(Stein))
	fmt.Println(Schere.Beats(Papier))
	fmt.Println()

	fmt.Println(Stein.Beats(Schere))
	fmt.Println(Stein.Beats(Stein))
	fmt.Println(Stein.Beats(Papier))
	fmt.Println()

	fmt.Println(Papier.Beats(Schere))
	fmt.Println(Papier.Beats(Stein))
	fmt.Println(Papier.Beats(Papier))

	// Output:
	// false
	// false
	// true
	//
	// true
	// false
	// false
	//
	// false
	// true
	// false
}
