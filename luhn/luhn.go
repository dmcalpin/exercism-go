// Package luhn has a Valid function
package luhn

import (
	"strings"
	"unicode"
)

// Valid does a luhn check and returns
// true if the number is valid.
func Valid(numStr string) bool {
	numStr = strings.ReplaceAll(numStr, " ", "")
	total := 0
	totalDigits := 0
	for _, n := range numStr {
		if !unicode.IsNumber(n) {
			if unicode.IsSpace(n) {
				continue
			}
			return false
		}

		num := int(n - '0')

		if len(numStr)%2 == totalDigits%2 {
			num = num * 2
			if num > 9 {
				num -= 9
			}
		}
		total += num
		totalDigits++

	}

	if totalDigits == 1 {
		return false
	}

	if total%10 == 0 {
		return true
	}

	return false
}
