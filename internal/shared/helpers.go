package shared

import (
	"regexp"
)

func delimitersForNumbers(delimiters string) string {
	if delimiters != "" {
		return delimiters
	}

	return "\\s"
}

func matchDelimiters(delimiters string, ch string) bool {
	if delimiters == "" {
		return false
	}

	matched, _ := regexp.MatchString(delimiters, ch)
	return matched
}
