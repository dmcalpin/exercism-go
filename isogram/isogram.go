// Package isogram has an IsIsogram function
package isogram

import (
	"unicode"
	"unicode/utf8"
)

// IsIsogram detects if a word has no
// duplicate letters. Space characters
// and dashes are allowed
func IsIsogram(word string) bool {
	// letterCache is a map of runes used
	// to keep track of which runes have
	// ben found in word
	letterCache := make(
		map[rune]bool,
		// creates a map of the proper size
		utf8.RuneCountInString(word),
	)

	// loop over the runes in a word
	for _, letter := range word {
		// ignore anything that isn't a letter
		if !unicode.IsLetter(letter) {
			continue
		}
		// if the letter is uppercase convert
		// it to lowercase for faster comparison
		if unicode.IsUpper(letter) {
			letter = unicode.ToLower(letter)
		}
		// if we found an existing letter in
		// the cache then this is not an isogram
		if letterCache[letter] {
			return false
		}
		// add the current letter to the cache
		letterCache[letter] = true
	}

	// if we've made it this far, word is
	// an isogram
	return true
}
