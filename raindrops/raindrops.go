// Package raindrops has a Convert function
package raindrops

import "strconv"

// declare the possible sounds to keep the
// code DRY
const (
	pling = "Pling"
	plang = "Plang"
	plong = "Plong"
)

// Convert takes a number and returns rain sounds
// This code could be shortened by using a string
// builder, which would allow me to reduce if
// statements, but that would slow it down and
// increase memory usage.
// Nested if statments reduce duplicate checks.
// mod3 can be followed by mod 5 and mod7 checks,
// but mod5 can only be followed by mod7 checks,
// and mod7 can not be followed by anything.
func Convert(input int) string {
	// identify which numbers input is divisible by
	mod3 := input%3 == 0
	mod5 := input%5 == 0
	mod7 := input%7 == 0

	if mod3 {
		if mod5 {
			if mod7 {
				return pling + plang + plong
			}
			return pling + plang
		}
		if mod7 {
			return pling + plong
		}
		return pling
	}
	if mod5 {
		if mod7 {
			return plang + plong
		}
		return plang
	}
	if mod7 {
		return plong
	}
	if input%2 == 0 {
		// Itoa is an efficient way to
		// convert an int to a string
		return strconv.Itoa(input)
	}

	return "1"
}
