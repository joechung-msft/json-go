package shared

import (
	"regexp"
)

func parseArray(array string) ArrayToken {
	const (
		Scanning = 1 + iota
		Elements
		Delimiter
		End
	)

	var (
		matched bool
		mode    = Scanning
		pos     int
		slice   string
		token   = Array{values: nil}
	)

	for pos < len(array) && mode != End {
		ch := string(array[pos])

		switch mode {
		case Scanning:
			matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch)

			if matched {
				pos++
			} else if ch == "[" {
				pos++
				mode = Elements
			} else {
				panic("Expected '['")
			}

		case Elements:
			matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch)

			if matched {
				pos++
			} else if ch == "]" {
				if len(token.values) > 0 {
					panic("Unexpected ','")
				}

				pos++
				mode = End
			} else {
				slice = array[pos:]
				valueToken := parseValue(slice, "[,\\]\\s]")
				token.values = append(token.values, valueToken)
				pos += valueToken.skip
				mode = Delimiter
			}

		case Delimiter:
			matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch)

			if matched {
				pos++
			} else if ch == "]" {
				pos++
				mode = End
			} else if ch == "," {
				pos++
				mode = Elements
			} else {
				panic("Expected ',' or ']'")
			}

		case End:

		default:
			panic("Unrecognized mode")
		}
	}

	return ArrayToken{skip: pos, token: token}
}
