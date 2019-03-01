package accumulate

func Accumulate(stringray []string, operation func(string) string) []string {
	// stringray isn't a fish!
	// its a string-array
	var output []string
	// let's loop through each word in string array.
	for _, word := range stringray {
		// we can do the operation(word),
		// and then add the result to output.
		output = append(output, operation(word))
	}
	// return our nice result!
	return output
}
