# go-highlight
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dieperdev/go-highlight)](https://pkg.go.dev/github.com/dieperdev/go-highlight)

go-highlight can highlight substrings of text with customizable prefixes and suffixes in Go.

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

## License
The [MIT License](LICENSE)
