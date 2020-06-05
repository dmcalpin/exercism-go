// Package acronym handles functionality for shortening strings.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the first leatter of each word for the given string.
func Abbreviate(s string) string {
	var abbreviations strings.Builder
	// replace instances dash or underscore characters with a space
	s = regexp.MustCompile(`[-_]`).ReplaceAllString(s, " ")
	// split the string into slice of words based on spaces
	for _, word := range strings.Fields(s) {
		// write the byte to our abbreviation
		abbreviations.WriteByte(byte(word[0]))
	}
	// convert to uppercase and return string
	return strings.ToUpper(abbreviations.String())
}
