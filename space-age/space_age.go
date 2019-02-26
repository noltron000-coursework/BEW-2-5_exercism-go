package space

import (
	"errors"
	"fmt"
	"strings"
)

// type Planet struct {
// 	planet string
// }

func convertUnit(time float64, unit_in string, unit_out string) (float64, error) {
	unit_in = strings.ToLower(unit_in)
	unit_out = strings.ToLower(unit_out)
	var mult float64
	var divi float64
	var err error

	switch unit_in {
	case "seconds":
		divi = 1.0 * 12 * (365 / 12) * 24 * 60 * 60
	case "minutes":
		divi = 1.0 * 12 * (365 / 12) * 24 * 60
	case "hours":
		divi = 1.0 * 12 * (365 / 12) * 24
	case "days":
		divi = 1.0 * 12 * (365 / 12)
	case "months":
		divi = 1.0 * 12
	case "years":
		divi = 1.0
	default:
		err = errors.New("That isn't a valid unit!")
	}

	switch unit_out {
	case "seconds":
		mult = 1.0 * 12 * (365.25 / 12) * 24 * 60 * 60
	case "minutes":
		mult = 1.0 * 12 * (365.25 / 12) * 24 * 60
	case "hours":
		mult = 1.0 * 12 * (365.25 / 12) * 24
	case "days":
		mult = 1.0 * 12 * (365.25 / 12)
	case "months":
		mult = 1.0 * 12
	case "years":
		mult = 1.0
	default:
		err = errors.New("That isn't a valid unit!")
	}

	return time * mult / divi, err
}

func convertAge(age float64, planet string) (float64, error) {
	planet = strings.ToLower(planet)
	var mult float64
	var err error

	fmt.Println(planet)

	switch planet {
	case "earth":
		mult = 1.0
	case "mercury":
		mult = 0.2408467
	case "venus":
		mult = 0.61519726
	case "mars":
		mult = 1.8808158
	case "jupiter":
		mult = 11.862615
	case "saturn":
		mult = 29.447498
	case "uranus":
		mult = 84.016846
	case "neptune":
		mult = 164.79132
	default:
		err = errors.New("That's not a valid planet!")
	}

	return age / mult, err
}

func Age(seconds float64, planet string) float64 {
	// var new_unit, _ = convertUnit(seconds, "seconds", "years")
	// ↑this↑ is actually more accurate than ↓this↓
	var new_unit = seconds / 31557600
	var new_age, _ = convertAge(new_unit, planet)
	return new_age
}
