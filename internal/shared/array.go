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
		matched      bool
		mode         = Scanning
		pos          int
		slice        string
		token        = Array{Values: nil}
		whitespaceRE = regexp.MustCompile(`[ \n\r\t]`)
	)

	for pos < len(array) && mode != End {
		ch := string(array[pos])

		switch mode {
		case Scanning:
			if matched = whitespaceRE.MatchString(ch); matched {
				pos++
			} else if ch == "[" {
				pos++
				mode = Elements
			} else {
				panic("Expected '['")
			}

		case Elements:
			if matched = whitespaceRE.MatchString(ch); matched {
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
			if matched = whitespaceRE.MatchString(ch); matched {
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

	if mode != End {
		panic("Unterminated array")
	}

	return ArrayToken{Skip: pos, Token: token}
}
