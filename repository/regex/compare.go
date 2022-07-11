package regex

import (
	"regexp"
	"strings"
)

type RegexCompare struct {
	Text string
}

func (r RegexCompare) Compare(pattern string) bool {
	pattern = strings.TrimSpace(pattern)

	cases := []string{
		strings.ToLower(pattern),
		strings.ToUpper(pattern),
		strings.Title(strings.ToLower(pattern)),
	}

	for i := 0; i < len(cases); i++ {
		if regexp.MustCompile(cases[i]).Match([]byte(r.Text)) {
			return true
		}
	}

	return false
}
