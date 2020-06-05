// Package twofer has a ShareWith function
package twofer

// ShareWith returns string describing who we're
// sharing with
func ShareWith(name string) string {
	// if name is empty, set name to "you"
	if name == "" {
		name = "you"
	}
	// concat the strings and return
	return "One for " + name + ", one for me."
}
