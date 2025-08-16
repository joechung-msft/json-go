package shared

import (
	"testing"
)

func TestFalse(t *testing.T) {
	actual := Parse("false")

	switch interface{}(actual.Token).(type) {
	case FalseToken:

	default:
		t.Error("false should be a FalseToken")
	}
}

func TestNull(t *testing.T) {
	actual := Parse("null")

	switch interface{}(actual.Token).(type) {
	case NullToken:

	default:
		t.Error("null should be a NullToken")
	}
}

func TestTrue(t *testing.T) {
	actual := Parse("true")

	switch interface{}(actual.Token).(type) {
	case TrueToken:

	default:
		t.Error("true should be a TrueToken")
	}
}
