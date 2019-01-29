package raindrops

import (
	"strconv"
	"fmt"
)

func Convert(drops int) (string) {
	fmt.Println(drops)
	var sound string = ""
	if drops % 3 == 0 {
		sound += "Pling"
	}
	if drops % 5 == 0 {
		sound += "Plang"
	}
	if drops % 7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound += strconv.Itoa(drops)
	}
	return sound
}