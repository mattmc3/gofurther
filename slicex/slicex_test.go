package slicex_test

import (
	"reflect"
	"testing"

	"github.com/mattmc3/gofurther/slicex"
)

type IndexOfTestParams struct {
	s    []interface{}
	item interface{}
	out  int
}

func TestIndexOf(t *testing.T) {
	var testCases = []IndexOfTestParams{
		{[]interface{}{1, 2, 3, 4}, 3, 2},
		{[]interface{}{"1", "2", "3", "4"}, 3, -1},
		{[]interface{}{"a", "b", "b", "a"}, "a", 0},
		{[]interface{}{"a", "b", "b", "a"}, "b", 1},
		{[]interface{}{"a", "b", "b", "a"}, "A", -1},
	}

	for _, test := range testCases {
		actual := slicex.IndexOf(test.s, test.item)
		if actual != test.out {
			t.Errorf(`IndexOf("%v",%v) = %v; want %v`, test.s, test.item, actual, test.out)
		}
	}
}

type ToStringsParams struct {
	in  []interface{}
	out []string
}

func TestStringify(t *testing.T) {
	var testCases = []ToStringsParams{
		{[]interface{}{}, []string{}},
		{[]interface{}{"1", 2, "a", 3.14, []rune{'g'}}, []string{"1", "2", "a", "3.14", "[103]"}},
	}

	for _, test := range testCases {
		actual := slicex.Stringify(test.in)
		if !reflect.DeepEqual(actual, test.out) {
			t.Errorf(`Stringify([]interface{}%v) = %v; want %v`, test.in, actual, test.out)
		}
	}
}
