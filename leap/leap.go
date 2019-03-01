// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// checks if the year is a leap year.
func IsLeapYear(year int) bool {
	//   year%4 == 0 -> year is divisible by 4! could be a leap year...
	//   && (etc...) -> if this 'etc' is false, its not a leap year.
	// year%100 != 0 -> year is NOT divisible 100. Good to go!
	// year%400 == 0 -> year is divisible by 400! definately a leap year!
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
