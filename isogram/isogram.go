// Package isogram has an IsIsogram function
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram detects if a word has no
// duplicate letters. Space characters
// and dashes are allowed
func IsIsogram(word string) bool {
	// loop over the runes in a word
	for i, letter := range word {
		// ignore anything that isn't a letter
		if !unicode.IsLetter(letter) {
			continue
		}
		// get a string of letters that come
		// after the current letter
		subWord := word[i+1:]
		// check if the letter exists
		if strings.ContainsRune(subWord, letter) {
			return false
		}
		if unicode.IsUpper(letter) {
			// if the letter is uppercase do a lower
			// case check
			if strings.ContainsRune(
				subWord,
				unicode.ToLower(letter),
			) {
				return false
			}
		} else {
			// if the letter is lowercase do a upeer
			// case check
			if strings.ContainsRune(
				subWord,
				unicode.ToUpper(letter),
			) {
				return false
			}
		}

	}

	// if we've made it this far, word is
	// an isogram
	return true
}
