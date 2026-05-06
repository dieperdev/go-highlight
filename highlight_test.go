package highlight

import (
"fmt"
	"testing"
)

type testCase struct {
	original      string
	substring     string
	limit         int
	before        string
	after         string
	caseSensitive bool

	testName string
	want     string
}

var testCases = []testCase{
	// Case-insensitive test cases
	testCase{
		original:      "hello world",
		substring:     "world",
		limit:         1,
		before:        "testing ",
		after:         ".",
		caseSensitive: false,

		testName: "\"hello world\", limit: 1, case-sensitive: false",
		want:     "hello testing world.",
	},
	testCase{
		original:      "hello world",
		substring:     "world",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: false,

		testName: "\"hello world\", limit: 2, case-sensitive: false",
		want:     "hello testing world.",
	},
	testCase{
		original:      "hello world world world",
		substring:     "world",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: false,

		testName: "\"hello world world world\", limit: 2, case-sensitive: false",
		want:     "hello testing world. testing world. world",
	},
	testCase{
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
	testCase{
		original:      "hello world",
		substring:     "World",
		limit:         1,
		before:        "testing ",
		after:         ".",
		caseSensitive: true,

		testName: "\"hello world\", limit: 1, case-sensitive: true",
		want:     "hello world",
	},
	testCase{
		original:      "hello World",
		substring:     "World",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: true,

		testName: "\"hello World\", limit: 2, case-sensitive: true",
		want:     "hello testing World.",
	},
	testCase{
		original:      "hello WORLD WORLD WORLD",
		substring:     "WORLD",
		limit:         2,
		before:        "testing ",
		after:         ".",
		caseSensitive: true,

		testName: "\"hello WORLD WORLD WORLD\", limit: 2, case-sensitive: true",
		want:     "hello testing WORLD. testing WORLD. WORLD",
	},
	testCase{
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

func TestHighlight(t *testing.T) {
	for _, tCase := range testCases {
		testName := fmt.Sprintf("testing %s", tCase.testName)

		t.Run(testName, func(t *testing.T) {
			highlighted := Highlight(tCase.original, tCase.substring, tCase.limit, tCase.before, tCase.after, tCase.caseSensitive)

			if highlighted != tCase.want {
				t.Errorf("got %q, want %q", highlighted, tCase.want)
			}
		})
	}
}
