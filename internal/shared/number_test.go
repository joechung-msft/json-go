package shared

import (
	"strconv"
	"testing"
)

func TestNumber(t *testing.T) {
	if !isNumberToken(Parse("0").Token) {
		t.Error("0 should be a NumberToken")
	}
}

func TestZeroes(t *testing.T) {
	zeroes := []string{
		"0",
		"-0",
		"0.",
		"-0.",
		"0.0",
		"-0.0",
		"0e0",
		"0e+0",
		"0e-0",
		"-0e0",
		"-0e+0",
		"-0e-0",
		"0.e0",
		"0.e+0",
		"0.e-0",
		"0.0e0",
		"0.0e+0",
		"0.0e-0",
		"-0.e0",
		"-0.e+0",
		"-0.e-0",
		"-0.0e0",
		"-0.0e+0",
		"-0.0e-0",
	}
	for _, zero := range zeroes {
		token := Parse(zero).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", zero)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.ValueAsString != zero {
			t.Errorf("%s should equal %s", number.ValueAsString, zero)
		} else if number.Value != 0 {
			t.Errorf("%s should equal 0", zero)
		}
	}
}

func isNumberToken(token any) bool {
	switch any(token).(type) {
	case NumberToken:
		return true
	default:
		return false
	}
}

func TestPositiveIntegers(t *testing.T) {
	integers := []string{
		"1",
		"42",
		"123456789",
		"999999999999999",
	}
	for _, num := range integers {
		token := Parse(num).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", num)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.ValueAsString != num {
			t.Errorf("%s should equal %s", number.ValueAsString, num)
		}
		expected, _ := strconv.ParseFloat(num, 64)
		if number.Value != expected {
			t.Errorf("%s should equal %f", num, expected)
		}
	}
}

func TestNegativeIntegers(t *testing.T) {
	integers := []string{
		"-1",
		"-42",
		"-123456789",
		"-999999999999999",
	}
	for _, num := range integers {
		token := Parse(num).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", num)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.ValueAsString != num {
			t.Errorf("%s should equal %s", number.ValueAsString, num)
		}
		expected, _ := strconv.ParseFloat(num, 64)
		if number.Value != expected {
			t.Errorf("%s should equal %f", num, expected)
		}
	}
}

func TestPositiveDecimals(t *testing.T) {
	decimals := []string{
		"0.1",
		"1.23",
		"42.0",
		"123.456",
		"999999999999999.999999999999999",
	}
	for _, num := range decimals {
		token := Parse(num).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", num)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.ValueAsString != num {
			t.Errorf("%s should equal %s", number.ValueAsString, num)
		}
		expected, _ := strconv.ParseFloat(num, 64)
		if number.Value != expected {
			t.Errorf("%s should equal %f", num, expected)
		}
	}
}

func TestNegativeDecimals(t *testing.T) {
	decimals := []string{
		"-0.1",
		"-1.23",
		"-42.0",
		"-123.456",
		"-999999999999999.999999999999999",
	}
	for _, num := range decimals {
		token := Parse(num).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", num)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.ValueAsString != num {
			t.Errorf("%s should equal %s", number.ValueAsString, num)
		}
		expected, _ := strconv.ParseFloat(num, 64)
		if number.Value != expected {
			t.Errorf("%s should equal %f", num, expected)
		}
	}
}

func TestNumbersWithExponents(t *testing.T) {
	exponents := []struct {
		input    string
		expected string
	}{
		{"1e10", "1e10"},
		{"1E10", "1e10"},
		{"1e+10", "1e+10"},
		{"1E+10", "1e+10"},
		{"1e-10", "1e-10"},
		{"1E-10", "1e-10"},
		{"-1e10", "-1e10"},
		{"-1E10", "-1e10"},
		{"-1e+10", "-1e+10"},
		{"-1E+10", "-1e+10"},
		{"-1e-10", "-1e-10"},
		{"-1E-10", "-1e-10"},
		{"123.456e78", "123.456e78"},
		{"123.456E78", "123.456e78"},
		{"123.456e+78", "123.456e+78"},
		{"123.456E+78", "123.456e+78"},
		{"123.456e-78", "123.456e-78"},
		{"123.456E-78", "123.456e-78"},
		{"-123.456e78", "-123.456e78"},
		{"-123.456E78", "-123.456e78"},
		{"-123.456e+78", "-123.456e+78"},
		{"-123.456E+78", "-123.456e+78"},
		{"-123.456e-78", "-123.456e-78"},
		{"-123.456E-78", "-123.456e-78"},
	}
	for _, test := range exponents {
		token := Parse(test.input).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", test.input)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.ValueAsString != test.expected {
			t.Errorf("%s should equal %s", number.ValueAsString, test.expected)
		}
		expected, _ := strconv.ParseFloat(test.input, 64)
		if number.Value != expected {
			t.Errorf("%s should equal %f", test.input, expected)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	edgeCases := []string{
		"1.7976931348623157e+308", // Max float64
		"2.2250738585072014e-308", // Min positive float64
		"9007199254740991",        // Max safe integer
		"-9007199254740991",       // Min safe integer
		"0.0000000000000001",      // Very small decimal
	}
	for _, num := range edgeCases {
		token := Parse(num).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", num)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.ValueAsString != num {
			t.Errorf("%s should equal %s", number.ValueAsString, num)
		}
		expected, _ := strconv.ParseFloat(num, 64)
		if number.Value != expected {
			t.Errorf("%s should equal %f", num, expected)
		}
	}
}

func TestNumbersWithWhitespace(t *testing.T) {
	whitespaceCases := []string{
		" 42",
		"42 ",
		" 42 ",
		"\t42",
		"\n42",
		"\r42",
		" \t\n\r42 \t\n\r",
	}
	for _, num := range whitespaceCases {
		token := Parse(num).Token
		if !isNumberToken(token) {
			t.Errorf("%s should be a NumberToken", num)
		}

		numberToken := token.(NumberToken)
		number := numberToken.Token
		if number.Value != 42 {
			t.Errorf("%s should parse to 42", num)
		}
	}
}

func TestInvalidNumbers(t *testing.T) {
	invalidCases := []string{
		"01",    // Leading zero
		"00",    // Leading zero
		"0.0.0", // Multiple decimal points
		"1e",    // Incomplete exponent
		"1e+",   // Incomplete exponent
		"1e-",   // Incomplete exponent
		"1ee10", // Double exponent
		"1.2.3", // Multiple decimal points
		"+1",    // Leading plus (not allowed in JSON)
		"1.",    // Trailing decimal point
		".1",    // Leading decimal point
		"1e1.0", // Decimal in exponent
		"1e1e1", // Double exponent
		"",      // Empty string
		"abc",   // Non-numeric
	}

	for _, num := range invalidCases {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("%s should have caused a panic", num)
			}
		}()
		Parse(num)
	}
}
