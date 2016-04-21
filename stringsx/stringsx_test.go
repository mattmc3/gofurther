package stringsx_test

import (
	"testing"

	"github.com/mattmc3/gofurther/stringsx"
)

type SnipTestParams struct {
	s        string
	startIdx int
	endIdx   int
	out      string
}

type SnipToEndTestParams struct {
	s        string
	startIdx int
	out      string
}

func TestSnip(t *testing.T) {
	var testCases = []SnipTestParams{
		{"abcd", 0, 4, "abcd"},
		{"abcd", 1, 5, "bcd"},
		{"abcd", -3, -1, "bc"},
		{"abcd", -999, -2, "ab"},
		{"abcd", 2, 2, ""},
		{"abcd", 997, 999, ""},
	}

	for _, test := range testCases {
		actual := stringsx.Snip(test.s, test.startIdx, test.endIdx)
		if actual != test.out {
			t.Errorf(`Snip("%v",%v,%v) = "%v"; want "%v"`, test.s, test.startIdx, test.endIdx, actual, test.out)
		}
	}
}

func TestSnipToEnd(t *testing.T) {
	var testCases = []SnipToEndTestParams{
		{"abcd", 0, "abcd"},
		{"abcd", 1, "bcd"},
		{"abcd", -3, "bcd"},
		{"abcd", -999, "abcd"},
		{"abcd", 3, "d"},
		{"abcd", 4, ""},
		{"abcd", 999, ""},
	}

	for _, test := range testCases {
		actual := stringsx.SnipToEnd(test.s, test.startIdx)
		if actual != test.out {
			t.Errorf(`SnipToEnd("%v",%v) = "%v"; want "%v"`, test.s, test.startIdx, actual, test.out)
		}
	}
}

var aligntests = []struct {
	in      string
	padWith rune
	length  int
	left    string
	right   string
	center  string
}{
	{"abc", ' ', -1, "abc", "abc", "abc"},
	{"abc", ' ', 0, "abc", "abc", "abc"},
	{"abc", ' ', 5, "abc  ", "  abc", " abc "},
	{"abc", '-', 6, "abc---", "---abc", "-abc--"},
}

func TestAlign(t *testing.T) {
	for _, tt := range aligntests {
		fn := func(name string, align func(in string, pad rune, l int) string, expected string) {
			actual := align(tt.in, tt.padWith, tt.length)
			if actual != expected {
				t.Errorf(`%v("%v", '%v', %v) = "%v"; want "%v"`, name, tt.in, string(tt.padWith), tt.length, actual, expected)
			}
		}
		fn("AlignLeft", stringsx.AlignLeft, tt.left)
		fn("AlignRight", stringsx.AlignRight, tt.right)
		fn("AlignCenter", stringsx.AlignCenter, tt.center)
	}
}
