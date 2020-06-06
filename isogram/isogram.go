// Package isogram has a IsIsogram function
package isogram

import "unicode"

// IsIsogram detects if a word has no
// duplicate letters. Space characters
// and dashes are allowed
func IsIsogram(word string) bool {
	letterCache := make(map[rune]bool)

	if word == "" {
		return true
	}

	for _, letter := range word {
		if letter == '-' || letter == ' ' {
			continue
		}
		normalizedLetter := unicode.ToLower(letter)
		if letterCache[normalizedLetter] {
			return false
		}
		letterCache[normalizedLetter] = true
	}

	return true
}
