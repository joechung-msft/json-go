package shared

import (
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

func isNumberToken(token interface{}) bool {
	switch interface{}(token).(type) {
	case NumberToken:
		return true
	default:
		return false
	}
}
