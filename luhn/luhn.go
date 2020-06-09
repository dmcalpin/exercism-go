// Package luhn has a Valid function
package luhn

import (
	"unicode"
)

// Valid does a luhn check and returns
// true if the number is valid.
func Valid(numStr string) bool {
	// convert string to rune slice
	numRunes := []rune(numStr)
	// keep track of total instead of storing
	// values in separate slice
	total := 0
	// since the numStr could contain spaces
	// just keep track of what digit we're on
	// instead of removing spaces up front
	currDigit := 0
	// loop over the rune starting from the end
	for i := len(numRunes) - 1; i >= 0; i-- {
		n := numRunes[i]
		if !unicode.IsNumber(n) {
			// spaces are allowed so continue
			if unicode.IsSpace(n) {
				continue
			}
			// anything else fails the check
			return false
		}

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
	}

	// the number is invalid if there's only
	// 1 digit or if the total is not divisible
	// by 10
	if currDigit == 1 || total%10 != 0 {
		return false
	}

	// if we got this far it's valid!
	return true
}
