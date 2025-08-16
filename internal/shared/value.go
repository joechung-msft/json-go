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
		matched bool
		mode    = Scanning
		pos     int
		slice   string
		token   interface{}
	)

	for pos < len(value) && mode != End {
		ch := string(value[pos])

		switch mode {
		case Scanning:
			if matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch); matched {
				pos++
			} else if ch == "[" {
				mode = Array
			} else if ch == "f" {
				mode = False
			} else if ch == "n" {
				mode = Null
			} else if matched, _ = regexp.MatchString("[-\\d]", ch); matched {
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
				panic("Unexpected character")
			}

		case Array:
			slice = value[pos:]
			arrayToken := parseArray(slice)
			token = arrayToken
			pos += arrayToken.skip
			mode = End

		case False:
			slice = value[pos : pos+5]
			if slice == "false" {
				token = FalseToken{value: false}
				pos += 5
				mode = End
			} else {
				panic("Expected 'false'")
			}

		case Null:
			slice = value[pos : pos+4]
			if slice == "null" {
				token = NullToken{value: nil}
				pos += 4
				mode = End
			} else {
				panic("Expected 'null'")
			}

		case Number:
			slice = value[pos:]
			numberToken := parseNumber(slice, delimitersForNumbers(delimiters))
			token = numberToken
			pos += numberToken.skip
			mode = End

		case Object:
			slice = value[pos:]
			objectToken := parseObject(slice)
			token = objectToken
			pos += objectToken.skip
			mode = End

		case String:
			slice = value[pos:]
			stringToken := parseString(slice)
			token = stringToken
			pos += stringToken.skip
			mode = End

		case True:
			slice = value[pos : pos+4]
			if slice == "true" {
				token = TrueToken{value: true}
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

	return ValueToken{skip: pos, token: token}
}
