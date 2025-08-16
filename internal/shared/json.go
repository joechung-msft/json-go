package shared

import (
	"regexp"
)

// Parse JSON
func Parse(json string) ValueToken {
	const (
		Scanning = 1 + iota
		Value
		End
	)

	var (
		matched bool
		mode    = Scanning
		pos     int
		slice   string
		token   ValueToken
	)

	for pos < len(json) && mode != End {
		ch := string(json[pos])

		switch mode {
		case Scanning:
			if matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch); matched {
				pos++
			} else {
				mode = Value
			}

		case Value:
			slice = json[pos:]
			token = parseValue(slice, "")
			pos += token.Skip
			mode = End

		case End:

		default:
			panic("Unrecognized mode")
		}
	}

	return token
}
