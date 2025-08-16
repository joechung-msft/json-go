package shared

import (
	"regexp"
)

func parsePair(pair string, delimiters string) PairToken {
	const (
		Scanning = 1 + iota
		String
		Delimiter
		Value
		End
	)

	var (
		matched bool
		mode    = Scanning
		pos     int
		slice   string
		token   Pair
	)

	for pos < len(pair) && mode != End {
		ch := string(pair[pos])

		switch mode {
		case Scanning:
			if matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch); matched {
				pos++
			} else {
				mode = String
			}

		case String:
			slice = pair[pos:]
			stringToken := parseString(slice)
			token.Key = stringToken
			pos += stringToken.Skip
			mode = Delimiter

		case Delimiter:
			if matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch); matched {
				pos++
			} else if ch == ":" {
				pos++
				mode = Value
			} else {
				panic("Expected ':'")
			}

		case Value:
			slice = pair[pos:]
			valueToken := parseValue(slice, delimiters)
			token.Value = valueToken
			pos += valueToken.Skip
			mode = End

		case End:

		default:
			panic("Unrecognized mode")
		}
	}

	return PairToken{Skip: pos, Token: token}
}
