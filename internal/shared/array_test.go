package shared

import (
	"testing"
)

func TestArray(t *testing.T) {
	actual := Parse("[]")

	switch any(actual.Token).(type) {
	case ArrayToken:

	default:
		t.Error("[] should be an Array")
	}
}

func TestDanglingCommaInArray(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("[1,] should have caused a panic")
		}
	}()

	Parse("[1,]")
}

func TestInvalidArray(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("[,] should have caused a panic")
		}
	}()

	Parse("[,]")
}
