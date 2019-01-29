package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) (bool) {
	word = strings.ToLower(word)
	var used string = ""
	var isogram bool = true

	// first for loop checks each item in the given word
	for i := 0; i < len(word); i++ {
		// second for loop checks already iterated items
		for j := 0; j < len(used); j++ {
			if word[i] == used[j] {
				isogram = false
				return isogram
			}
		}
		if unicode.IsLetter(rune(word[i])) {
			used += string(word[i])
		}
	}
	return isogram
}