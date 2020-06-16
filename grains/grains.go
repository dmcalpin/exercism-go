package grains

// BadIndexError is returned when an invalid
// board index is used
type BadIndexError struct{}

func (e BadIndexError) Error() string {
	return "Index must be between 1 and 64"
}

// Square takes an index and returns the
// total square of grains
func Square(index int) (uint64, error) {
	// If the index is out of range
	// return an error
	if index <= 0 || index >= 65 {
		return 0, BadIndexError{}
	}
	// shifting a bit one to the
	// left squares the number
	// Binary => Decimal
	// 1 => 1
	// 10 => 2
	// 100 => 4
	// 1000 => 8
	return 1 << (index - 1), nil
}

// Total returns the total number of grains
func Total() uint64 {
	return 18446744073709551615
}
