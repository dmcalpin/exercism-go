// Package luhn has a Valid function
package luhn

import (
	"unicode"
)

// Valid does a luhn check and returns
// true if the number is valid.
func Valid(numStr string) bool {
	// keep track of total instead of storing
	// values in separate slice
	total := 0
	// since the numStr could contain spaces
	// just keep track of what digit we're on
	// instead of removing spaces up front
	currDigit := 0
	// loop over the rune starting from the end
	for i := len(numStr) - 1; i >= 0; i-- {
		n := rune(numStr[i])
		switch {
		case unicode.IsNumber(n):
			// convert the rune to it's int value
			num := int(n - '0')

			// if we're on an odd digit do some
			// luhn logic
			if currDigit%2 == 1 {
				num = num * 2
				if num > 9 {
					num -= 9
				}
			}
			// increase the running total
			total += num
			// increase the current digit
			currDigit++
		case unicode.IsSpace(n):
			// spaces are allowed so continue
			continue
		default:
			// anything else fails the check
			return false
		}
	}

	// the number is valid if there's more than
	// one digit and it's divisible by 10
	return currDigit > 1 && total%10 == 0
}
