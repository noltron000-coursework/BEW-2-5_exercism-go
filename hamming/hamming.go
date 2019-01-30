package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	// hamming num starts at 0
	var hamming int = 0
	// start with no error
	var err error = nil

	// ensure a & b are of equal lengths
	if len(a) == len(b) {
		// index (i) starts at 0, iterate through len(a), adding one each time.
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hamming += 1
			}
		}
		return hamming, err

		// throw an error -- a & b are different lengths
	} else {
		err = errors.New("DNA sequences have uneven length!")
		return hamming, err
	}
}
