package regexpx_test

import (
	"testing"

	"github.com/mattmc3/gofurther/regexpx"
)

type WildcardConvTestParams struct {
	wildcard string
	out      string
}

func TestConvertWildcardToRegexpr(t *testing.T) {
	var testCases = []WildcardConvTestParams{
		{"", ""},
		{"*.go|*.a", "^(.*\\.go)|(.*\\.a)$"},
	}

	for _, test := range testCases {
		actual := regexpx.ConvertWildcardToRegexpr(test.wildcard)
		if actual != test.out {
			t.Errorf(`ConvertWildcardToRegexpr("%v") = "%v"; want "%v"`, test.wildcard, actual, test.out)
		}
	}
}
