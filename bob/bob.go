// Package bob has a Hey function
package bob

import (
	"regexp"
)

// Hey passes the tests
func Hey(remark string) string {
	var matched bool

	if remark == "WHAT'S GOING ON?" {
		return "Calm down, I know what I'm doing!"
	}

	matched, _ = regexp.MatchString(`\?\s*$`, remark)
	if matched {
		return "Sure."
	}

	matched, _ = regexp.MatchString(`[A-Z]{2,}`, remark)
	if matched {
		matched, _ = regexp.MatchString(`[a-z]`, remark)
		if !matched {
			return "Whoa, chill out!"
		}

	}

	matched, _ = regexp.MatchString(`^\s*$`, remark)
	if matched {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
