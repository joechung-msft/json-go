package shared

import (
	"regexp"
)

var whitespaceRegex = regexp.MustCompile(`[ \n\r\t]`)

func parseObject(object string) ObjectToken {
	const (
		Scanning = 1 + iota
		Pair
		Delimiter
		End
	)

	var (
		mode  = Scanning
		pos   int
		slice string
		token = Object{Members: nil}
	)

	for pos < len(object) && mode != End {
		ch := string(object[pos])

		switch mode {
		case Scanning:
			if whitespaceRegex.MatchString(ch) {
				pos++
			} else if ch == "{" {
				pos++
				mode = Pair
			} else {
				panic("Expected '{'")
			}

		case Pair:
			if whitespaceRegex.MatchString(ch) {
				pos++
			} else if ch == "}" {
				if len(token.Members) > 0 {
					panic("Unexpected ','")
				}

				pos++
				mode = End
			} else {
				slice = object[pos:]
				pairToken := parsePair(slice, "[\\s,\\}]")
				token.Members = append(token.Members, pairToken)
				pos += pairToken.Skip
				mode = Delimiter
			}

		case Delimiter:
			if whitespaceRegex.MatchString(ch) {
				pos++
			} else if ch == "," {
				pos++
				mode = Pair
			} else if ch == "}" {
				pos++
				mode = End
			} else {
				panic("Expected ',' or '}'")
			}

		case End:

		default:
			panic("Unrecognized mode")
		}
	}

	return ObjectToken{Skip: pos, Token: token}
}
