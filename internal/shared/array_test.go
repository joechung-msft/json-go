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

func TestArrayWithSingleString(t *testing.T) {
	actual := Parse(`["hello"]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 1 {
			t.Errorf("Expected 1 element, got %d", len(token.Token.Values))
		}
		if str, ok := token.Token.Values[0].Token.(StringToken); ok {
			if str.Token != "hello" {
				t.Errorf("Expected 'hello', got '%s'", str.Token)
			}
		} else {
			t.Error("Expected string token")
		}
	default:
		t.Error(`["hello"] should be an Array`)
	}
}

func TestArrayWithSingleNumber(t *testing.T) {
	actual := Parse(`[42]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 1 {
			t.Errorf("Expected 1 element, got %d", len(token.Token.Values))
		}
		if num, ok := token.Token.Values[0].Token.(NumberToken); ok {
			if num.Token.Value != 42 {
				t.Errorf("Expected 42, got %f", num.Token.Value)
			}
		} else {
			t.Error("Expected number token")
		}
	default:
		t.Error(`[42] should be an Array`)
	}
}

func TestArrayWithSingleBoolean(t *testing.T) {
	actual := Parse(`[true]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 1 {
			t.Errorf("Expected 1 element, got %d", len(token.Token.Values))
		}
		if boolToken, ok := token.Token.Values[0].Token.(TrueToken); ok {
			if !boolToken.Value {
				t.Error("Expected true")
			}
		} else {
			t.Error("Expected true token")
		}
	default:
		t.Error(`[true] should be an Array`)
	}
}

func TestArrayWithSingleNull(t *testing.T) {
	actual := Parse(`[null]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 1 {
			t.Errorf("Expected 1 element, got %d", len(token.Token.Values))
		}
		if _, ok := token.Token.Values[0].Token.(NullToken); !ok {
			t.Error("Expected null token")
		}
	default:
		t.Error(`[null] should be an Array`)
	}
}

func TestArrayWithMultipleElements(t *testing.T) {
	actual := Parse(`["string", 123, true, null]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 4 {
			t.Errorf("Expected 4 elements, got %d", len(token.Token.Values))
		}

		// Check string
		if str, ok := token.Token.Values[0].Token.(StringToken); ok {
			if str.Token != "string" {
				t.Errorf("Expected 'string', got '%s'", str.Token)
			}
		} else {
			t.Error("Expected string token at index 0")
		}

		// Check number
		if num, ok := token.Token.Values[1].Token.(NumberToken); ok {
			if num.Token.Value != 123 {
				t.Errorf("Expected 123, got %f", num.Token.Value)
			}
		} else {
			t.Error("Expected number token at index 1")
		}

		// Check boolean
		if boolToken, ok := token.Token.Values[2].Token.(TrueToken); ok {
			if !boolToken.Value {
				t.Error("Expected true at index 2")
			}
		} else {
			t.Error("Expected true token at index 2")
		}

		// Check null
		if _, ok := token.Token.Values[3].Token.(NullToken); !ok {
			t.Error("Expected null token at index 3")
		}
	default:
		t.Error(`["string", 123, true, null] should be an Array`)
	}
}

func TestArrayWithWhitespace(t *testing.T) {
	actual := Parse(`[ "hello" , 42 ]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(token.Token.Values))
		}
	default:
		t.Error(`[ "hello" , 42 ] should be an Array`)
	}
}

func TestNestedArrays(t *testing.T) {
	actual := Parse(`[[1, 2], [3, 4]]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(token.Token.Values))
		}

		// Check first nested array
		if arr1, ok := token.Token.Values[0].Token.(ArrayToken); ok {
			if len(arr1.Token.Values) != 2 {
				t.Errorf("Expected 2 elements in first nested array, got %d", len(arr1.Token.Values))
			}
		} else {
			t.Error("Expected array token at index 0")
		}

		// Check second nested array
		if arr2, ok := token.Token.Values[1].Token.(ArrayToken); ok {
			if len(arr2.Token.Values) != 2 {
				t.Errorf("Expected 2 elements in second nested array, got %d", len(arr2.Token.Values))
			}
		} else {
			t.Error("Expected array token at index 1")
		}
	default:
		t.Error(`[[1, 2], [3, 4]] should be an Array`)
	}
}

func TestArrayWithObjects(t *testing.T) {
	actual := Parse(`[{"key": "value"}, {"number": 123}]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 2 {
			t.Errorf("Expected 2 elements, got %d", len(token.Token.Values))
		}

		// Check first object
		if obj1, ok := token.Token.Values[0].Token.(ObjectToken); ok {
			if len(obj1.Token.Members) != 1 {
				t.Errorf("Expected 1 member in first object, got %d", len(obj1.Token.Members))
			}
		} else {
			t.Error("Expected object token at index 0")
		}

		// Check second object
		if obj2, ok := token.Token.Values[1].Token.(ObjectToken); ok {
			if len(obj2.Token.Members) != 1 {
				t.Errorf("Expected 1 member in second object, got %d", len(obj2.Token.Members))
			}
		} else {
			t.Error("Expected object token at index 1")
		}
	default:
		t.Error(`[{"key": "value"}, {"number": 123}] should be an Array`)
	}
}

func TestArrayWithMixedTypes(t *testing.T) {
	actual := Parse(`["string", 42, true, false, null, {"key": "value"}, [1, 2, 3]]`)

	switch token := any(actual.Token).(type) {
	case ArrayToken:
		if len(token.Token.Values) != 7 {
			t.Errorf("Expected 7 elements, got %d", len(token.Token.Values))
		}
	default:
		t.Error("Mixed types array should be parsed correctly")
	}
}

func TestArrayTrailingCommaPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("[1, 2,] should have caused a panic")
		}
	}()

	Parse("[1, 2,]")
}

func TestArrayMissingCommaPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("[1 2] should have caused a panic")
		}
	}()

	Parse("[1 2]")
}

func TestArrayUnclosedPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("[1, 2 should have caused a panic")
		}
	}()

	Parse("[1, 2")
}
