// Package diffsquares provides functions
// for identifying sums of squares,
// squares of sums and finding the
// difference
package diffsquares

// SquareOfSum adds numbers and then
// squares them. This can be expressed as
// the mathematical equation: [n(n/2 + 0.5)]^2
func SquareOfSum(num int) (total int) {
	sqrt := int(((float64(num) * 0.5) + 0.5) *
		float64(num))
	return sqrt * sqrt
}

// SumOfSquares squares the given number,
// and all numbers before it, and then adds
// them together. This can be expressed as the
// mathematical equation: Σn2 = [n(n+1)(2n+1)]/6
func SumOfSquares(num int) (total int) {
	return (num * (num + 1) * (2*num + 1)) / 6
}

// Difference subtracts the SumofSquares
// from the SquareOfSum
func Difference(num int) (diff int) {
	return SquareOfSum(num) - SumOfSquares(num)
}
