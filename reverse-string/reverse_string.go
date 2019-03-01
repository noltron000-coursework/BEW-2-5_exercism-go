package reverse

func String(gnirts string) string {
	// gnirts is a backwards string !
	// uh-oh! lets reverse it :)
	var output string
	for _, letter := range gnirts {
		output = string(letter) + output
	}
	return output
}
