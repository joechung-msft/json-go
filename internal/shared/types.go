package shared

// Array is a JSON array
type Array struct {
	values []ValueToken
}

// ArrayToken is a JSON array token
type ArrayToken struct {
	skip  int
	token Array
}

// FalseToken is a JSON Boolean false token
type FalseToken struct {
	value bool
}

// NullToken is a JSON null token
type NullToken struct {
	value interface{}
}

// Number is a JSON number
type Number struct {
	value         float64
	valueAsString string
}

// NumberToken is a JSON number token
type NumberToken struct {
	skip  int
	token Number
}

// Object is a JSON object
type Object struct {
	members []PairToken
}

// ObjectToken is a JSON object token
type ObjectToken struct {
	skip  int
	token Object
}

// Pair is a JSON pair
type Pair struct {
	key   StringToken
	value ValueToken
}

// PairToken is a JSON pair token
type PairToken struct {
	skip  int
	token Pair
}

// StringToken is a JSON string token
type StringToken struct {
	skip  int
	token string
}

// TrueToken is a JSON Boolean true token
type TrueToken struct {
	value bool
}

// ValueToken is a JSON value token
type ValueToken struct {
	skip  int
	token interface{}
}
