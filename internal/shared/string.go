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
		matched bool
		mode    = Scanning
		pos     int
		slice   string
		token   string
	)

	for pos < len(s) && mode != End {
		ch := string(s[pos])

		switch mode {
		case Scanning:
			if matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch); matched {
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
			if ch == "\"" || ch == "\\" || ch == "/" {
				token += ch
				pos++
				mode = Character
			} else if ch == "b" {
				token += "\b"
				pos++
				mode = Character
			} else if ch == "f" {
				token += "\f"
				pos++
				mode = Character
			} else if ch == "n" {
				token += "\n"
				pos++
				mode = Character
			} else if ch == "r" {
				token += "\r"
				pos++
				mode = Character
			} else if ch == "t" {
				token += "\t"
				pos++
				mode = Character
			} else if ch == "u" {
				pos++
				mode = Unicode
			} else {
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

	return StringToken{skip: pos, token: token}
}
