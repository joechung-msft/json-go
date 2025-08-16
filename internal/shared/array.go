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
		token   = Array{Values: nil}
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
				if len(token.Values) > 0 {
					panic("Unexpected ','")
				}

				pos++
				mode = End
			} else {
				slice = array[pos:]
				valueToken := parseValue(slice, "[,\\]\\s]")
				token.Values = append(token.Values, valueToken)
				pos += valueToken.Skip
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

	return ArrayToken{Skip: pos, Token: token}
}
