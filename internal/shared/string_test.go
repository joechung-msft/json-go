package shared

import (
	"testing"
)

func TestString(t *testing.T) {
	actual := Parse("\"Hello, world!\"")

	switch interface{}(actual.token).(type) {
	case StringToken:

	default:
		t.Error("\"Hello, world!\" should be a StringToken")
	}
}
