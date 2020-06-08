// Package luhn has a Valid function
package luhn

import (
	"unicode"
)

// Valid does a luhn check and returns
// true if the number is valid.
func Valid(num string) bool {

	nums := []int{}
	for _, n := range num {
		if !unicode.IsNumber(n) {
			if unicode.IsSpace(n) {
				continue
			}
			return false
		}
		nums = append(nums, int(n-'0'))
	}

	if len(nums) == 1 {
		return false
	}

	for i := len(nums) - 2; i >= 0; i -= 2 {
		n := nums[i] * 2
		if n > 9 {
			n -= 9
		}
		nums[i] = n
	}

	total := 0
	for _, n := range nums {
		total += n
	}

	if total%10 == 0 {
		return true
	}

	return false
}
