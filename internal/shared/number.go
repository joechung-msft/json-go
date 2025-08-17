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
		mode           = Scanning
		pos            int
		token          = Number{Value: math.NaN(), ValueAsString: ""}
		whitespaceRe   = regexp.MustCompile(`[ \n\r\t]`)
		digitRe        = regexp.MustCompile(`\d`)
		nonZeroDigitRe = regexp.MustCompile(`[1-9]`)
		exponentRe     = regexp.MustCompile(`[eE]`)
	)

	for pos < len(number) && mode != End {
		ch := string(number[pos])

		switch mode {
		case Scanning:
			if whitespaceRe.MatchString(ch) {
				pos++
			} else if ch == "-" {
				pos++
				token.ValueAsString += "-"
			}

			mode = Characteristic

		case Characteristic:
			if ch == "0" {
				pos++
				token.ValueAsString += "0"
				mode = DecimalPoint
			} else if nonZeroDigitRe.MatchString(ch) {
				pos++
				token.ValueAsString += ch
				mode = CharacteristicDigit
			} else {
				panic("Expected digit")
			}

		case CharacteristicDigit:
			if digitRe.MatchString(ch) {
				pos++
				token.ValueAsString += ch
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				mode = DecimalPoint
			}

		case DecimalPoint:
			if ch == "." {
				pos++
				token.ValueAsString += "."
				mode = Mantissa
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				mode = Exponent
			}

		case Mantissa:
			if digitRe.MatchString(ch) {
				pos++
				token.ValueAsString += ch
			} else if exponentRe.MatchString(ch) {
				mode = Exponent
			} else if matchDelimiters(delimiters, ch) {
				mode = End
			} else {
				panic("Unexpected character")
			}

		case Exponent:
			if ch == "e" || ch == "E" {
				pos++
				token.ValueAsString += "e"
				mode = ExponentSign
			} else {
				panic("Expected 'e' or 'E'")
			}

		case ExponentSign:
			if ch == "+" || ch == "-" {
				pos++
				token.ValueAsString += ch
			}

			mode = ExponentFirstDigit

		case ExponentFirstDigit:
			if digitRe.MatchString(ch) {
				pos++
				token.ValueAsString += ch
				mode = ExponentDigits
			} else {
				panic("Expected digit")
			}

		case ExponentDigits:
			if digitRe.MatchString(ch) {
				pos++
				token.ValueAsString += ch
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
		if value, err := strconv.ParseFloat(token.ValueAsString, 64); err != nil {
			panic(err)
		} else {
			token.Value = value
		}
	}

	return NumberToken{Skip: pos, Token: token}
}
