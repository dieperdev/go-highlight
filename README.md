# go-highlight
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dieperdev/go-highlight)](https://pkg.go.dev/github.com/dieperdev/go-highlight)

go-highlight highlights substrings of text with customizable prefixes and suffixes in Go.

## Examples

```go
import (
        "fmt"

        "github.com/dieperdev/go-highlight"
)

fmt.Println(highlight.Highlight("Hello World!", "World", 1, "GitHub ", "", false))
// Output: Hello GitHub World!



// Colors can't be shown on the README.
// You will have to run an example locally like the one
// under `highlight_example_test.go`

fmt.Println(highlight.HighlightColor("Red, Yellow, and Blue.", "Yellow", 1, "Orange", true))
// Output: Red, OrangeYellowOrange, and Blue.



fmt.Println(highlight.HighlightSameReplacements("Go is a programming language.", "Go", 1, "*", true))
// Output: *Go* is a programming language.
}
```

## Using highlighted text
If you want to print the highlighted text to the console, you need a library such as [k0kubun/go-ansi](https://github.com/k0kubun/go-ansi) to do it.
`go-ansi` will correctly print the ANSI color escape codes (it also has Windows support).
You should avoid using the `%q` verb (quoted string) when `fmt.Sprintf`'ing or manipulating the highlighted text as it will escape the highlights and not print correctly.

## License
The [MIT License](LICENSE)
