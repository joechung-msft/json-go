package shared

// Array is a JSON array
type Array struct {
	Values []ValueToken
}

// ArrayToken is a JSON array token
type ArrayToken struct {
	Skip  int
	Token Array
}

// FalseToken is a JSON Boolean false token
type FalseToken struct {
	Value bool
}

// NullToken is a JSON null token
type NullToken struct {
	Value interface{}
}

// Number is a JSON number
type Number struct {
	Value         float64
	ValueAsString string
}

// NumberToken is a JSON number token
type NumberToken struct {
	Skip  int
	Token Number
}

// Object is a JSON object
type Object struct {
	Members []PairToken
}

// ObjectToken is a JSON object token
type ObjectToken struct {
	Skip  int
	Token Object
}

// Pair is a JSON pair
type Pair struct {
	Key   StringToken
	Value ValueToken
}

// PairToken is a JSON pair token
type PairToken struct {
	Skip  int
	Token Pair
}

// StringToken is a JSON string token
type StringToken struct {
	Skip  int
	Token string
}

// TrueToken is a JSON Boolean true token
type TrueToken struct {
	Value bool
}

// ValueToken is a JSON value token
type ValueToken struct {
	Skip  int
	Token interface{}
}
