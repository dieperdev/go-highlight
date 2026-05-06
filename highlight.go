package highlight

import (
	gostring "github.com/boyter/go-string"
)

// Highlight prefixes up to `limit` occurrences of `substring`
// with `before` and suffixes up to `limit` occurrences of
// `substring` with `after`.
//
// Highlight can optionally be case-sensitive.
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

// HighlightColor has the same functionality as Highlight, but prefixes and
// suffixes up to `limit` occurences of `substring` with the same text.
//
// The `color` passed to HighlightColor does not have to strictly be a color.
// You may want to use `HighlightSameReplacements` to make usage of
// `go-highlight` more clear in your program if `color` is not a color.
func HighlightColor(original string, substring string, limit int, color string, caseSensitive bool) string {
	return Highlight(original, substring, limit, color, color, caseSensitive)
}

// HighlightSameReplacements has the same functionality as HighlightColor, but
// with different function name. It is intended to make usage of `go-highlight`
// more clear.
func HighlightSameReplacements(original string, substring string, limit int, replacement string, caseSensitive bool) string {
	return Highlight(original, substring, limit, replacement, replacement, caseSensitive)
}
