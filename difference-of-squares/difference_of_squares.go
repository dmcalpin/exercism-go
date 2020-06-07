// Package diffsquares provides functions
// for identifying sums of squares,
// squares of sums and finding the
// difference
package diffsquares

// SquareOfSum adds numbers and then
// squares them
func SquareOfSum(num int) (total int) {
	// while num is positive
	for num > 0 {
		// add the num to the total
		total += num
		// and count down by 1
		num--
	}
	// return the square of the total
	return total * total
}

// SumOfSquares squares numbers and
// then addes them
func SumOfSquares(num int) (total int) {
	// while num is positive
	for num > 0 {
		// add the square of the num to
		// the total
		total += num * num
		// and countdown by 1
		num--
	}
	// return the total
	return
}

// Difference subtracts the SumofSquares
// from the SquareOfSum
func Difference(num int) (diff int) {
	// simply take the difference of the two
	// functions
	return SquareOfSum(num) - SumOfSquares(num)
}
