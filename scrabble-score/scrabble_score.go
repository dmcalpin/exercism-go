// Package scrabble has a Score function
package scrabble

// Score takes a word and returns the
// points according to Scrabble rules.
func Score(word string) int {
	// score starts at 0
	totalScore := 0

	for _, letter := range word {
		// for each letter in the word, add to
		// the total score. Both upper and lower
		// case are included to eliminate the
		// need to call unicode.ToUpper
		switch letter {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T',
			'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			totalScore++
		case 'D', 'G', 'd', 'g':
			totalScore += 2
		case 'B', 'C', 'M', 'P', 'b', 'c', 'm', 'p':
			totalScore += 3
		case 'F', 'H', 'V', 'W', 'Y', 'f', 'h', 'v', 'w', 'y':
			totalScore += 4
		case 'K', 'k':
			totalScore += 5
		case 'J', 'X', 'j', 'x':
			totalScore += 8
		case 'Q', 'Z', 'q', 'z':
			totalScore += 10
		}
	}

	return totalScore
}
