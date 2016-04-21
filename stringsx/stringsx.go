package stringsx

import "strings"

// Snip allows for pulling apart pieces of a string by indexes. Indexes need
// not be within the range of characters in the string.
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

// AlignRight returns a string that right-aligns the characters of the string
// provided by padding them on the left with ahe specified rune to the specified
// total length.
// s: The string to pad
// padChar: The character to use for padding
// toLength: The desired minimum length of the resulting string. If the input
//   string length exceeds this value, the original value is returned.
func AlignRight(s string, padChar rune, toLength int) string {
	return getPadding(s, padChar, toLength) + s
}

// AlignLeft returns a string that left-aligns the characters of the string
// provided by padding them on the right with the specified rune to the
// specified total length.
// s: The string to pad
// padChar: The character to use for padding
// toLength: The desired minimum length of the resulting string. If the input
//   string length exceeds this value, the original value is returned.
func AlignLeft(s string, padChar rune, toLength int) string {
	return s + getPadding(s, padChar, toLength)
}

// AlignCenter returns a string that center-aligns the characters of the string
// provided by padding them on both sides with the specified rune to the
// specified total length.
// s: The string to pad
// padChar: The character to use for padding
// toLength: The desired minimum length of the resulting string. If the input
//   string length exceeds this value, the original value is returned.
func AlignCenter(s string, padChar rune, toLength int) string {
	var padding = getPadding(s, padChar, toLength)
	l := len(padding)
	return padding[:l/2] + s + padding[l/2:]
}

func getPadding(s string, padChar rune, totalWidth int) string {
	if len(s) >= totalWidth {
		return ""
	}
	var padding = strings.Repeat(string(padChar), totalWidth-len(s))
	return padding
}
