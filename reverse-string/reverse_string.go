// Package reverse has a Reverse function
package reverse

// Reverse reverses the order of letters
// for a given string
func Reverse(str string) string {
	// convert str to slice of runes
	runes := []rune(str)
	runeCount := len(runes)
	// only loop halfway through the slice
	// so that we don't undo our reverse
	for i := 0; i < runeCount/2; i++ {
		// swap the runes in the slice
		runes[i], runes[runeCount-i-1] =
			runes[runeCount-i-1], runes[i]
	}
	// conver the slice back to a string
	return string(runes)
}
