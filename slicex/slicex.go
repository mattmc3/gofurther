package slicex

import "fmt"

// IndexOf returns the index in the slice of the specified item. Returns -1 if
// the item is not found.
// s: the slice
// item: the item to find
func IndexOf(s []interface{}, item interface{}) int {
	for idx, val := range s {
		if val == item {
			return idx
		}
	}
	return -1
}

// Byteify takes any slice and converts the contents to their byte
// representation
func Byteify(s []interface{}) [][]byte {
	result := make([][]byte, len(s))
	for idx, val := range s {
		var bval []byte
		switch theval := val.(type) {
		case []byte:
			bval = theval
		case string:
			bval = []byte(theval)
		case fmt.Stringer:
			bval = []byte(theval.String())
		default:
			bval = []byte(fmt.Sprintf("%v", val))
		}
		result[idx] = bval
	}
	return result
}

// Stringify takes any slice and converts the contents to their string
// representation
func Stringify(s []interface{}) []string {
	result := make([]string, len(s))
	for idx, val := range s {
		var sval string
		switch str := val.(type) {
		case string:
			sval = str
		case fmt.Stringer:
			sval = str.String()
		default:
			sval = fmt.Sprintf("%v", val)
		}
		result[idx] = sval
	}
	return result
}

// ObjectifyStrings takes a slice of strings and converts the contents to an
// []interface{} representation
func ObjectifyStrings(s []string) []interface{} {
	if s == nil {
		return nil
	}
	r := make([]interface{}, len(s))
	for i := range s {
		r[i] = s[i]
	}
	return r
}
