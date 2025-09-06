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

func TestSimpleObject(t *testing.T) {
	actual := Parse("{\"key\": \"value\"}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(token.Token.Members))
		}
		pair := token.Token.Members[0]
		if pair.Token.Key.Token != "key" {
			t.Errorf("Expected key 'key', got %q", pair.Token.Key.Token)
		}
		if strVal, ok := pair.Token.Value.Token.(StringToken); ok {
			if strVal.Token != "value" {
				t.Errorf("Expected value 'value', got %q", strVal.Token)
			}
		} else {
			t.Error("Expected string value")
		}
	default:
		t.Error("{\"key\": \"value\"} should be an ObjectToken")
	}
}

func TestObjectMultiplePairs(t *testing.T) {
	actual := Parse("{\"a\": \"1\", \"b\": \"2\"}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 2 {
			t.Errorf("Expected 2 members, got %d", len(token.Token.Members))
		}
		// Check first pair
		pair1 := token.Token.Members[0]
		if pair1.Token.Key.Token != "a" {
			t.Errorf("Expected key 'a', got %q", pair1.Token.Key.Token)
		}
		if strVal, ok := pair1.Token.Value.Token.(StringToken); ok {
			if strVal.Token != "1" {
				t.Errorf("Expected value '1', got %q", strVal.Token)
			}
		} else {
			t.Error("Expected string value for 'a'")
		}
		// Check second pair
		pair2 := token.Token.Members[1]
		if pair2.Token.Key.Token != "b" {
			t.Errorf("Expected key 'b', got %q", pair2.Token.Key.Token)
		}
		if strVal, ok := pair2.Token.Value.Token.(StringToken); ok {
			if strVal.Token != "2" {
				t.Errorf("Expected value '2', got %q", strVal.Token)
			}
		} else {
			t.Error("Expected string value for 'b'")
		}
	default:
		t.Error("{\"a\": \"1\", \"b\": \"2\"} should be an ObjectToken")
	}
}

func TestObjectWithNumberValue(t *testing.T) {
	actual := Parse("{\"num\": 42}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(token.Token.Members))
		}
		pair := token.Token.Members[0]
		if pair.Token.Key.Token != "num" {
			t.Errorf("Expected key 'num', got %q", pair.Token.Key.Token)
		}
		if numVal, ok := pair.Token.Value.Token.(NumberToken); ok {
			if numVal.Token.Value != 42 {
				t.Errorf("Expected value 42, got %f", numVal.Token.Value)
			}
		} else {
			t.Error("Expected number value")
		}
	default:
		t.Error("{\"num\": 42} should be an ObjectToken")
	}
}

func TestObjectWithBooleanValues(t *testing.T) {
	actual := Parse("{\"trueVal\": true, \"falseVal\": false}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 2 {
			t.Errorf("Expected 2 members, got %d", len(token.Token.Members))
		}
		// Check true value
		pair1 := token.Token.Members[0]
		if pair1.Token.Key.Token != "trueVal" {
			t.Errorf("Expected key 'trueVal', got %q", pair1.Token.Key.Token)
		}
		if boolVal, ok := pair1.Token.Value.Token.(TrueToken); ok {
			if !boolVal.Value {
				t.Error("Expected true value")
			}
		} else {
			t.Error("Expected true value")
		}
		// Check false value
		pair2 := token.Token.Members[1]
		if pair2.Token.Key.Token != "falseVal" {
			t.Errorf("Expected key 'falseVal', got %q", pair2.Token.Key.Token)
		}
		if boolVal, ok := pair2.Token.Value.Token.(FalseToken); ok {
			if boolVal.Value {
				t.Error("Expected false value")
			}
		} else {
			t.Error("Expected false value")
		}
	default:
		t.Error("{\"trueVal\": true, \"falseVal\": false} should be an ObjectToken")
	}
}

func TestObjectWithNullValue(t *testing.T) {
	actual := Parse("{\"nullVal\": null}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(token.Token.Members))
		}
		pair := token.Token.Members[0]
		if pair.Token.Key.Token != "nullVal" {
			t.Errorf("Expected key 'nullVal', got %q", pair.Token.Key.Token)
		}
		if _, ok := pair.Token.Value.Token.(NullToken); !ok {
			t.Error("Expected null value")
		}
	default:
		t.Error("{\"nullVal\": null} should be an ObjectToken")
	}
}

func TestNestedObject(t *testing.T) {
	actual := Parse("{\"outer\": {\"inner\": \"value\"}}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(token.Token.Members))
		}
		pair := token.Token.Members[0]
		if pair.Token.Key.Token != "outer" {
			t.Errorf("Expected key 'outer', got %q", pair.Token.Key.Token)
		}
		if objVal, ok := pair.Token.Value.Token.(ObjectToken); ok {
			if len(objVal.Token.Members) != 1 {
				t.Errorf("Expected 1 inner member, got %d", len(objVal.Token.Members))
			}
			innerPair := objVal.Token.Members[0]
			if innerPair.Token.Key.Token != "inner" {
				t.Errorf("Expected inner key 'inner', got %q", innerPair.Token.Key.Token)
			}
			if strVal, ok := innerPair.Token.Value.Token.(StringToken); ok {
				if strVal.Token != "value" {
					t.Errorf("Expected inner value 'value', got %q", strVal.Token)
				}
			} else {
				t.Error("Expected string inner value")
			}
		} else {
			t.Error("Expected object value")
		}
	default:
		t.Error("{\"outer\": {\"inner\": \"value\"}} should be an ObjectToken")
	}
}

func TestObjectWithArrayValue(t *testing.T) {
	actual := Parse("{\"arr\": [1, 2, 3]}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(token.Token.Members))
		}
		pair := token.Token.Members[0]
		if pair.Token.Key.Token != "arr" {
			t.Errorf("Expected key 'arr', got %q", pair.Token.Key.Token)
		}
		if arrVal, ok := pair.Token.Value.Token.(ArrayToken); ok {
			if len(arrVal.Token.Values) != 3 {
				t.Errorf("Expected 3 array elements, got %d", len(arrVal.Token.Values))
			}
			// Check first element
			if numVal, ok := arrVal.Token.Values[0].Token.(NumberToken); ok {
				if numVal.Token.Value != 1 {
					t.Errorf("Expected first element 1, got %f", numVal.Token.Value)
				}
			} else {
				t.Error("Expected number in array")
			}
		} else {
			t.Error("Expected array value")
		}
	default:
		t.Error("{\"arr\": [1, 2, 3]} should be an ObjectToken")
	}
}

func TestObjectWithWhitespace(t *testing.T) {
	actual := Parse("{ \"key\" : \"value\" }")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(token.Token.Members))
		}
		pair := token.Token.Members[0]
		if pair.Token.Key.Token != "key" {
			t.Errorf("Expected key 'key', got %q", pair.Token.Key.Token)
		}
		if strVal, ok := pair.Token.Value.Token.(StringToken); ok {
			if strVal.Token != "value" {
				t.Errorf("Expected value 'value', got %q", strVal.Token)
			}
		} else {
			t.Error("Expected string value")
		}
	default:
		t.Error("{ \"key\" : \"value\" } should be an ObjectToken")
	}
}

func TestObjectWithEscapedString(t *testing.T) {
	actual := Parse("{\"key\\n\": \"value\\\"with\\\"quotes\"}")

	switch token := any(actual.Token).(type) {
	case ObjectToken:
		if len(token.Token.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(token.Token.Members))
		}
		pair := token.Token.Members[0]
		if pair.Token.Key.Token != "key\n" {
			t.Errorf("Expected key 'key\\n', got %q", pair.Token.Key.Token)
		}
		if strVal, ok := pair.Token.Value.Token.(StringToken); ok {
			expected := "value\"with\"quotes"
			if strVal.Token != expected {
				t.Errorf("Expected value %q, got %q", expected, strVal.Token)
			}
		} else {
			t.Error("Expected string value")
		}
	default:
		t.Error("{\"key\\n\": \"value\\\"with\\\"quotes\"} should be an ObjectToken")
	}
}
