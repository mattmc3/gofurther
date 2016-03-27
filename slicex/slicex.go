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

// Stringify takes any slice and converts the contents to their string
// representation
func Stringify(s []interface{}) []string {
	result := make([]string, len(s))
	for idx, val := range s {
		result[idx] = fmt.Sprintf("%v", val)
	}
	return result
}
