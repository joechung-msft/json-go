package shared

import (
	"regexp"
	"strconv"
)

func parseString(s string) StringToken {
	const (
		Scanning = 1 + iota
		Character
		EscapedCharacter
		Unicode
		End
	)

	var (
		mode        = Scanning
		pos         int
		slice       string
		token       string
		spaceRegexp = regexp.MustCompile(`[ \n\r\t]`)
	)

	for pos < len(s) && mode != End {
		ch := string(s[pos])

		switch mode {
		case Scanning:
			if spaceRegexp.MatchString(ch) {
				pos++
			} else if ch == "\"" {
				pos++
				mode = Character
			} else {
				panic("Expected '\"'")
			}

		case Character:
			if ch == "\\" {
				pos++
				mode = EscapedCharacter
			} else if ch == "\"" {
				pos++
				mode = End
			} else if ch != "\n" && ch != "\r" {
				token += ch
				pos++
			} else {
				panic("Unexpected character")
			}

		case EscapedCharacter:
			switch ch {
			case "\"", "\\", "/":
				token += ch
				pos++
				mode = Character
			case "b":
				token += "\b"
				pos++
				mode = Character
			case "f":
				token += "\f"
				pos++
				mode = Character
			case "n":
				token += "\n"
				pos++
				mode = Character
			case "r":
				token += "\r"
				pos++
				mode = Character
			case "t":
				token += "\t"
				pos++
				mode = Character
			case "u":
				pos++
				mode = Unicode
			default:
				panic("Unexpected escape character")
			}

		case Unicode:
			slice = s[pos : pos+4]
			if hex, err := strconv.ParseInt(slice, 16, 32); err == nil {
				token += string(rune(hex))
				pos += 4
				mode = Character
			} else {
				panic("Unrecognized Unicode code")
			}

		case End:

		default:
			panic("Unrecognized mode")
		}
	}

	if mode != End {
		panic("Unterminated string")
	}

	return StringToken{Skip: pos, Token: token}
}
