package highlight

import (
	"fmt"
	"strings"
	"testing"
)

type highlightTestCase struct {
	original      string
	substring     string
	limit         int
	before        string
	after         string
	caseSensitive bool

	testName string
	want     string
}

type highlightColorTestCase struct {
	original      string
	substring     string
	limit         int
	color         string
	caseSensitive bool

	testName string
	want     string
}

type highlightSameReplacementsTestCase struct {
	original      string
	substring     string
	limit         int
	replacement   string
	caseSensitive bool

	testName string
	want     string
}

var highlightTestCases = []highlightTestCase{
	// Case-insensitive test cases
	highlightTestCase{
		original:      "hello world",
		substring:     "world",
		limit:         1,
		before:        "testing ",
		after:         ".",
		caseSensitive: false,

		testName: "\"hello world\", limit: 1, case-sensitive: false",
		want:     "hello testing world.",
	},
	highlightTestCase{
		original:      "hello world",
		substring:     "world",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: false,

		testName: "\"hello world\", limit: 2, case-sensitive: false",
		want:     "hello testing world.",
	},
	highlightTestCase{
		original:      "hello world world world",
		substring:     "world",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: false,

		testName: "\"hello world world world\", limit: 2, case-sensitive: false",
		want:     "hello testing world. testing world. world",
	},
	highlightTestCase{
		original:      "hello world world world",
		substring:     "world",
		limit:         3,
		before:        "testing ",
		after:         ".",
		caseSensitive: false,

		testName: "\"hello world world world\", limit: 3, case-sensitive: false",
		want:     "hello testing world. testing world. testing world.",
	},

	// Case-sensitive test cases
	highlightTestCase{
		original:      "hello world",
		substring:     "World",
		limit:         1,
		before:        "testing ",
		after:         ".",
		caseSensitive: true,

		testName: "\"hello world\", limit: 1, case-sensitive: true",
		want:     "hello world",
	},
	highlightTestCase{
		original:      "hello World",
		substring:     "World",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: true,

		testName: "\"hello World\", limit: 2, case-sensitive: true",
		want:     "hello testing World.",
	},
	highlightTestCase{
		original:      "hello WORLD WORLD WORLD",
		substring:     "WORLD",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: true,

		testName: "\"hello WORLD WORLD WORLD\", limit: 2, case-sensitive: true",
		want:     "hello testing WORLD. testing WORLD. WORLD",
	},
	highlightTestCase{
		original:      "hello woRld woRld woRld",
		substring:     "world",
		limit:         3,
		before:        "testing ",
		after:         ".",
		caseSensitive: true,

		testName: "\"hello woRld woRld woRld\", limit: 3, case-sensitive: true",
		want:     "hello woRld woRld woRld",
	},
}

var highlightColorTestCases = []highlightColorTestCase{
	// Case-insensitive test cases
	highlightColorTestCase{
		original:      "hello world",
		substring:     "world",
		limit:         1,
		color:         FgBlack,
		caseSensitive: false,

		testName: "\"hello world\", limit: 1, case-sensitive: false",
		want:     fmt.Sprintf("hello %sworld%s", FgBlack, FgBlack),
	},
	highlightColorTestCase{
		original:      "hello world",
		substring:     "world",
		limit:         2,
		color:         FgBlack,
		caseSensitive: false,

		testName: "\"hello world\", limit: 2, case-sensitive: false",
		want:     fmt.Sprintf("hello %sworld%s", FgBlack, FgBlack),
	},
	highlightColorTestCase{
		original:      "hello world world world",
		substring:     "world",
		limit:         2,
		color:         FgBlack,
		caseSensitive: false,

		testName: "\"hello world world world\", limit: 2, case-sensitive: false",
		want:     fmt.Sprintf("hello %sworld%s %sworld%s world", FgBlack, FgBlack, FgBlack, FgBlack),
	},
	highlightColorTestCase{
		original:      "hello world world world",
		substring:     "world",
		limit:         3,
		color:         FgBlack,
		caseSensitive: false,

		testName: "\"hello world world world\", limit: 3, case-sensitive: false",
		want:     fmt.Sprintf("hello %sworld%s %sworld%s %sworld%s", FgBlack, FgBlack, FgBlack, FgBlack, FgBlack, FgBlack),
	},

	// Case-sensitive test cases
	highlightColorTestCase{
		original:      "hello world",
		substring:     "World",
		limit:         1,
		color:         FgBlack,
		caseSensitive: true,

		testName: "\"hello world\", limit: 1, case-sensitive: true",
		want:     "hello world",
	},
	highlightColorTestCase{
		original:      "hello World",
		substring:     "World",
		limit:         2,
		color:         FgBlack,
		caseSensitive: true,

		testName: "\"hello World\", limit: 2, case-sensitive: true",
		want:     fmt.Sprintf("hello %sWorld%s", FgBlack, FgBlack),
	},
	highlightColorTestCase{
		original:      "hello WORLD WORLD WORLD",
		substring:     "WORLD",
		limit:         2,
		color:         FgBlack,
		caseSensitive: true,

		testName: "\"hello WORLD WORLD WORLD\", limit: 2, case-sensitive: true",
		want:     fmt.Sprintf("hello %sWORLD%s %sWORLD%s WORLD", FgBlack, FgBlack, FgBlack, FgBlack),
	},
	highlightColorTestCase{
		original:      "hello woRld woRld woRld",
		substring:     "world",
		limit:         3,
		color:         FgBlack,
		caseSensitive: true,

		testName: "\"hello woRld woRld woRld\", limit: 3, case-sensitive: true",
		want:     "hello woRld woRld woRld",
	},
}

var highlightSameReplacementsTestCases = make([]highlightSameReplacementsTestCase, len(highlightColorTestCases))

// Fill the `SameReplacements` test cases with the same data as the `Color`
// test cases but change the field name to `replacement` and change the
// replacement string to "new".
func init() {
	for i, tCase := range highlightColorTestCases {
		highlightSameReplacementsTestCases[i] = highlightSameReplacementsTestCase{
			original:      tCase.original,
			substring:     tCase.substring,
			limit:         tCase.limit,
			replacement:   strings.ReplaceAll(tCase.color, FgBlack, "new"),
			caseSensitive: tCase.caseSensitive,

			testName: tCase.testName,
			want:     strings.ReplaceAll(tCase.want, FgBlack, "new"),
		}
	}
}

func TestHighlight(t *testing.T) {
	for _, tCase := range highlightTestCases {
		testName := fmt.Sprintf("testing %s", tCase.testName)

		t.Run(testName, func(t *testing.T) {
			highlighted := Highlight(tCase.original, tCase.substring, tCase.limit, tCase.before, tCase.after, tCase.caseSensitive)

			if highlighted != tCase.want {
				t.Errorf("got %q, want %q", highlighted, tCase.want)
			}
		})
	}
}

func TestHighlightColor(t *testing.T) {
	for _, tCase := range highlightColorTestCases {
		testName := fmt.Sprintf("testing %s", tCase.testName)

		t.Run(testName, func(t *testing.T) {
			highlighted := HighlightColor(tCase.original, tCase.substring, tCase.limit, tCase.color, tCase.caseSensitive)

			if highlighted != tCase.want {
				t.Errorf("got %q, want %q", highlighted, tCase.want)
			}
		})
	}
}

func TestHighlightSameReplacements(t *testing.T) {
	for _, tCase := range highlightSameReplacementsTestCases {
		testName := fmt.Sprintf("testing %s", tCase.testName)

		t.Run(testName, func(t *testing.T) {
			highlighted := HighlightSameReplacements(tCase.original, tCase.substring, tCase.limit, tCase.replacement, tCase.caseSensitive)

			if highlighted != tCase.want {
				t.Errorf("got %q, want %q", highlighted, tCase.want)
			}
		})
	}
}
