package highlight

import (
	gostring "github.com/boyter/go-string"
)

func Highlight(original string, substring string, limit int, before string, after string, caseSensitive bool) string {
	var positions [][]int

	if caseSensitive {
		positions = gostring.IndexAll(original, substring, limit)
	} else {
		positions = gostring.IndexAllIgnoreCase(original, substring, limit)
	}

	highlighted := gostring.HighlightString(original, positions, before, after)

	return highlighted
}
