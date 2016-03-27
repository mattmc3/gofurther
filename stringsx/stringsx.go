package stringsx

// Snip allows for pulling apart pieces of a string by indexes. Indexes need
// not be within the range of characters in th string.
// s: The string to act on
// startIndex: The starting index. Negative values move back from the end of the
// string. Indexes past the end of the string are allowed.
// endIndex: The ending index. Negative values move back from the end of the
// string. Indexes past the end of the string are
func Snip(s string, startIndex, endIndex int) string {
	strlen := len(s)

	fnRelToAbs := func(i int) int {
		result := i
		// negatives go from the right of the string
		if result < 0 {
			result = strlen + result
		}

		if result < 0 {
			result = 0
		} else if result > strlen {
			result = strlen
		}
		return result
	}

	startIndex = fnRelToAbs(startIndex)
	endIndex = fnRelToAbs(endIndex)
	if startIndex >= endIndex {
		return ""
	} else {
		return s[startIndex:endIndex]
	}
}

// SnipToEnd allows for pulling apart pieces of a string by indexes. Indexes need
// not be within the range of characters in the string.
// s: The string to act on
// startIndex: The starting index.
func SnipToEnd(s string, startIndex int) string {
	return Snip(s, startIndex, len(s))
}
