package shared

import (
	"math"
	"regexp"
	"strconv"
)

func parseNumber(number string, delimiters string) NumberToken {
	const (
		Scanning = 1 + iota
		Characteristic
		CharacteristicDigit
		DecimalPoint
		Mantissa
		Exponent
		ExponentSign
		ExponentFirstDigit
		ExponentDigits
		End
	)

	var (
		matched bool
		mode    = Scanning
		pos     int
		token   = Number{value: math.NaN(), valueAsString: ""}
	)

	for pos < len(number) && mode != End {
		ch := string(number[pos])

		switch mode {
		case Scanning:
			if matched, _ = regexp.MatchString("[ \\n\\r\\t]", ch); matched {
				pos++
			} else if ch == "-" {
				pos++
				token.valueAsString += "-"
			}

			mode = Characteristic

		case Characteristic:
			if ch == "0" {
				pos++
				token.valueAsString += "0"
				mode = DecimalPoint
			} else if matched, _ = regexp.MatchString("[1-9]", ch); matched {
				pos++
				token.valueAsString += ch
				mode = CharacteristicDigit
			} else {
				panic("Expected digit")
			}

		case CharacteristicDigit:
			if matched, _ = regexp.MatchString("\\d", ch); matched {
				pos++
				token.valueAsString += ch
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				mode = DecimalPoint
			}

		case DecimalPoint:
			if ch == "." {
				pos++
				token.valueAsString += "."
				mode = Mantissa
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				mode = Exponent
			}

		case Mantissa:
			if matched, _ = regexp.MatchString("\\d", ch); matched {
				pos++
				token.valueAsString += ch
			} else if matched, _ = regexp.MatchString("[eE]", ch); matched {
				mode = Exponent
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				panic("Unexpected character")
			}

		case Exponent:
			if ch == "e" || ch == "E" {
				pos++
				token.valueAsString += "e"
				mode = ExponentSign
			} else {
				panic("Expected 'e' or 'E'")
			}

		case ExponentSign:
			if ch == "+" || ch == "-" {
				pos++
				token.valueAsString += ch
			}

			mode = ExponentFirstDigit

		case ExponentFirstDigit:
			if matched, _ = regexp.MatchString("\\d", ch); matched {
				pos++
				token.valueAsString += ch
				mode = ExponentDigits
			} else {
				panic("Expected digit")
			}

		case ExponentDigits:
			if matched, _ = regexp.MatchString("\\d", ch); matched {
				pos++
				token.valueAsString += ch
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				panic("Expected digit")
			}

		case End:

		default:
			panic("Unrecogized mode")
		}
	}

	if mode == Characteristic || mode == ExponentFirstDigit {
		panic("Incomplete expression")
	} else {
		if value, err := strconv.ParseFloat(token.valueAsString, 64); err != nil {
			panic(err)
		} else {
			token.value = value
		}
	}

	return NumberToken{skip: pos, token: token}
}
