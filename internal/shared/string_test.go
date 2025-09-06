package shared

import (
	"testing"
)

func TestString(t *testing.T) {
	actual := Parse("\"Hello, world!\"")

	switch any(actual.Token).(type) {
	case StringToken:

	default:
		t.Error("\"Hello, world!\" should be a StringToken")
	}
}

func TestStringEmpty(t *testing.T) {
	actual := Parse("\"\"")

	switch token := any(actual.Token).(type) {
	case StringToken:
		if token.Token != "" {
			t.Errorf("Expected empty string, got %q", token.Token)
		}
	default:
		t.Error("\"\" should be a StringToken")
	}
}

func TestStringWithSpaces(t *testing.T) {
	actual := Parse("\"hello world\"")

	switch token := any(actual.Token).(type) {
	case StringToken:
		if token.Token != "hello world" {
			t.Errorf("Expected \"hello world\", got %q", token.Token)
		}
	default:
		t.Error("\"hello world\" should be a StringToken")
	}
}

func TestStringWithEscapedQuotes(t *testing.T) {
	actual := Parse("\"He said \\\"hello\\\"\"")

	switch token := any(actual.Token).(type) {
	case StringToken:
		expected := "He said \"hello\""
		if token.Token != expected {
			t.Errorf("Expected %q, got %q", expected, token.Token)
		}
	default:
		t.Error("Escaped quotes string should be a StringToken")
	}
}

func TestStringWithBackslash(t *testing.T) {
	actual := Parse("\"path\\\\to\\\\file\"")

	switch token := any(actual.Token).(type) {
	case StringToken:
		expected := "path\\to\\file"
		if token.Token != expected {
			t.Errorf("Expected %q, got %q", expected, token.Token)
		}
	default:
		t.Error("Backslash string should be a StringToken")
	}
}

func TestStringWithUnicode(t *testing.T) {
	actual := Parse("\"\\u0041\"")

	switch token := any(actual.Token).(type) {
	case StringToken:
		expected := "A"
		if token.Token != expected {
			t.Errorf("Expected %q, got %q", expected, token.Token)
		}
	default:
		t.Error("Unicode string should be a StringToken")
	}
}

func TestStringWithAllEscapes(t *testing.T) {
	actual := Parse("\"\\\"\\\\\\/\\b\\f\\n\\r\\t\"")

	switch token := any(actual.Token).(type) {
	case StringToken:
		expected := "\"\\/\b\f\n\r\t"
		if token.Token != expected {
			t.Errorf("Expected %q, got %q", expected, token.Token)
		}
	default:
		t.Error("All escapes string should be a StringToken")
	}
}

func TestStringUnterminated(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Unterminated string should have caused a panic")
		}
	}()

	Parse("\"hello")
}

func TestStringInvalidEscape(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Invalid escape should have caused a panic")
		}
	}()

	Parse("\"\\z\"")
}

func TestStringInvalidUnicode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Invalid unicode should have caused a panic")
		}
	}()

	Parse("\"\\u123\"")
}

func TestStringWithNewline(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("String with newline should have caused a panic")
		}
	}()

	Parse("\"hello\nworld\"")
}
