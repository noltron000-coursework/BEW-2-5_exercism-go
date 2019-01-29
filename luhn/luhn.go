package luhn

import (
	"regexp"
	"strconv"
)

func Valid(stringify string) bool {
	var numberify int = 0
	var boolify bool = false

	var cleanify1 string = ""
	regnum, _ := regexp.Compile("[0-9 ]")

	for index := 0; index < len(stringify); index++ {
		if regnum.MatchString(string(stringify[index])) {
			cleanify1 += string(stringify[index])
		} else {
			return boolify
		}
	}

	var cleanify2 string = ""
	regnum, _ = regexp.Compile("[0-9]")

	for index := 0; index < len(cleanify1); index++ {
		if regnum.MatchString(string(cleanify1[index])) {
			cleanify2 += string(cleanify1[index])
		}
	}

	for index := 0; index < len(cleanify2); index++ {
		cleanindex, _ := strconv.Atoi(string(cleanify2[index]))

		if (len(cleanify2)-index)%2 == 0 {
			if cleanindex*2 > 9 {
				numberify += cleanindex*2 - 9
			} else {
				numberify += cleanindex * 2
			}
		} else {

			numberify += cleanindex
		}
	}

	if numberify%10 == 0 && len(cleanify2) > 1 {
		boolify = true
	} else {
		boolify = false
	}

	return boolify
}
