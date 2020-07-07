// Package robotname does robot stuff
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot is a robot
type Robot struct {
	name string
}

var possibleNames = make([]string, maxNames)

func init() {
	// Pre-build all the possible robot names
	n := 0
	for i := 1; i <= 26; i++ {
		for j := 1; j <= 26; j++ {
			for k := 1; k <= 1000; k++ {
				possibleNames[n] = fmt.Sprintf(
					"%s%s%03d",
					string(rune(i+64)),
					string(rune(j+64)),
					k-1,
				)
				n++
			}
		}
	}
}

// Name returns the robot's name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(possibleNames) == 0 {
			return "", errors.New("No robot names left :-(")
		}
		// pick a random robot name from our slice
		nameIndex := rand.Intn(len(possibleNames))
		// set it
		r.name = possibleNames[nameIndex]
		// remove it from the slice
		possibleNames = unorderedRemove(possibleNames, nameIndex)
	}

	return r.name, nil
}

// Reset resets the robot's name
func (r *Robot) Reset() {
	// set to "" so r.Name() knows
	// it can set a new name
	r.name = ""
}

// remove an element from a slice at a given
// but does not maintain order
func unorderedRemove(slice []string, i int) []string {
	slice[i] = slice[len(slice)-1]
	slice[len(slice)-1] = ""
	return slice[:len(slice)-1]
}
