// Package hamming has a Distance function
package hamming

// MismatchedLengthError is used when sequences
// are of a different length.
type MismatchedLengthError struct{}

// MismatchedLengthError implements the
// error interface
func (err MismatchedLengthError) Error() string {
	return "Unable to compare squences of different length"
}

// Distance calculates the hamming distance
func Distance(a, b string) (int, error) {
	// Hamming distance is only applicable to
	// sequences of the same length. Check
	// that first.
	if len(a) != len(b) {
		// returning a custom error here.
		// could have used errors.New(), but
		// this turned out to be faster, and
		// gives semantic meaning to the error
		return 0, MismatchedLengthError{}
	}
	// if the two sequences are the same
	// we can just return 0 instead of
	// comparing each character
	if a == b {
		return 0, nil
	}

	// count will keep track of the differences
	// in each pair of sequences
	count := 0
	for i := range a {
		// a[i] and b[i] are bytes being
		// compared, if they are different
		// increment the count
		if a[i] != b[i] {
			count++
		}
	}
	// return the total number of differences
	return count, nil
}
