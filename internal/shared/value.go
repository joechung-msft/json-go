package shared

import (
	"regexp"
)

func parseValue(value string, delimiters string) ValueToken {
	const (
		Scanning = 1 + iota
		Array
		False
		Null
		Number
		Object
		String
		True
		End
	)

	var (
		mode         = Scanning
		pos          int
		slice        string
		token        any
		spaceRegexp  = regexp.MustCompile(`[ \n\r\t]`)
		numberRegexp = regexp.MustCompile(`[-\d]`)
	)

	for pos < len(value) && mode != End {
		ch := string(value[pos])

		switch mode {
		case Scanning:
			if spaceRegexp.MatchString(ch) {
				pos++
			} else if ch == "[" {
				mode = Array
			} else if ch == "f" {
				mode = False
			} else if ch == "n" {
				mode = Null
			} else if numberRegexp.MatchString(ch) {
				mode = Number
			} else if ch == "{" {
				mode = Object
			} else if ch == "\"" {
				mode = String
			} else if ch == "t" {
				mode = True
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				panic("Unexpected character in value: " + ch)
			}

		case Array:
			slice = value[pos:]
			arrayToken := parseArray(slice)
			token = arrayToken
			pos += arrayToken.Skip
			mode = End

		case False:
			if pos+5 > len(value) {
				panic("Unexpected end of input when parsing 'false'")
			}
			slice = value[pos : pos+5]
			if slice == "false" {
				token = FalseToken{Value: false}
				pos += 5
				mode = End
			} else {
				panic("Expected 'false'")
			}

		case Null:
			if pos+4 > len(value) {
				panic("Unexpected end of input when parsing 'null'")
			}
			slice = value[pos : pos+4]
			if slice == "null" {
				token = NullToken{Value: nil}
				pos += 4
				mode = End
			} else {
				panic("Expected 'null'")
			}

		case Number:
			slice = value[pos:]
			numberToken := parseNumber(slice, delimitersForNumbers(delimiters))
			token = numberToken
			pos += numberToken.Skip
			mode = End

		case Object:
			slice = value[pos:]
			objectToken := parseObject(slice)
			token = objectToken
			pos += objectToken.Skip
			mode = End

		case String:
			slice = value[pos:]
			stringToken := parseString(slice)
			token = stringToken
			pos += stringToken.Skip
			mode = End

		case True:
			if pos+4 > len(value) {
				panic("Unexpected end of input when parsing 'true'")
			}
			slice = value[pos : pos+4]
			if slice == "true" {
				token = TrueToken{Value: true}
				pos += 4
				mode = End
			} else {
				panic("Expected 'true'")
			}

		case End:

		default:
			panic("Unrecognized mode")
		}
	}

	return ValueToken{Skip: pos, Token: token}
}
