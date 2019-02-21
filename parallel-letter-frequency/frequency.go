package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(input string) FreqMap {
	dict := FreqMap{}
	for _, key := range input {
		dict[key]++
	}
	return dict
}

// SUBMITTING EXERCISE EARLY TO SEE OTHER SOLUTIONS
