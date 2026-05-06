package highlight_test

import (
	"fmt"
	"github.com/dieperdev/go-highlight"
)

func ExampleHighlight() {
	fmt.Println(highlight.Highlight("Hello World!", "World", 1, "GitHub ", "", false))
	// Output: Hello GitHub World!
}

func ExampleHighlightColor() {
	// Colors can't be printed on the online Go documentation.
	// You will have to run an example locally.
	//
	// You can copy this code example below for a quick test:
	/*
		package main

		import (
			"fmt"

			"github.com/dieperdev/go-highlight"
		)

		func main() {
			text := "Red, Yellow, and Blue."
			highlighted := highlight.HighlightColor(text, "Yellow", 1, highlight.BgYellow, true)

			fmt.Println(highlighted)
		}
	*/

	fmt.Println(highlight.HighlightColor("Red, Yellow, and Blue.", "Yellow", 1, "Orange", true))
	// Output: Red, OrangeYellowOrange, and Blue.
}

func ExampleHighlightSameReplacements() {
	fmt.Println(highlight.HighlightSameReplacements("Go is a programming language.", "Go", 1, "*", true))
	// Output: *Go* is a programming language.
}
