package regexpx

import (
	"bytes"
	"regexp"
	"strings"
)

// ConvertWildcardToRegexpr takes a wildcard pattern and converts it to a regular
// expression pattern. A wildcard pattern is like "*.go" or
func ConvertWildcardToRegexpr(wildcard string) (expr string) {
	expr = ""
	if wildcard == "" {
		return
	}

	var buf bytes.Buffer
	firstPass := true
	buf.WriteString("^")

	parts := strings.Split(wildcard, "|")
	for _, p := range parts {
		rePattern := regexp.QuoteMeta(p)
		rePattern = strings.Replace(rePattern, "\\[!", "[^", -1)
		rePattern = strings.Replace(rePattern, "\\[", "[", -1)
		rePattern = strings.Replace(rePattern, "\\]", "]", -1)
		rePattern = strings.Replace(rePattern, "\\?", ".", -1)
		rePattern = strings.Replace(rePattern, "\\*", ".*", -1)
		rePattern = strings.Replace(rePattern, "\\#", "\\d", -1)

		if firstPass {
			firstPass = false
		} else {
			buf.WriteString("|")
		}
		buf.WriteString("(")
		buf.WriteString(rePattern)
		buf.WriteString(")")
	}

	buf.WriteString("$")
	expr = buf.String()
	return
}
