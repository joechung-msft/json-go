package shared

import (
	"testing"
)

func TestObject(t *testing.T) {
	actual := Parse("{}")

	switch any(actual.Token).(type) {
	case ObjectToken:

	default:
		t.Error("{} should be an ObjectToken")
	}
}

func TestDanglingCommaInObject(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("{\"foo\":\"bar\",} should have caused a panic")
		}
	}()

	Parse("{\"foo\":\"bar\",}")
}
