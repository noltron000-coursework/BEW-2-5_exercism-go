// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

func ShareWith(name string) string {

	if name != "" {
		// return statement for adding variable name
		var myString = "One for " + name + ", one for me."
		return myString
	} else {
		// default return statement
		return "One for you, one for me."
	}
}
