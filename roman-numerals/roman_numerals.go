package romannumerals

import (
	"errors"
)

var values_map = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func ToRomanNumeral(input int) (string, error) {

	if input > 3999 || input < 1 {
		return "", errors.New("only positive values allowed. the largest number is MMMCMXCIX (or 3,999)")
	}

	roman_numeral := ""

	calcNumber := func(ceil int, floor int) {
		if input >= ceil {
			roman_numeral += values_map[ceil]
			input -= ceil
		} else {
			roman_numeral += values_map[floor]
			input -= floor
		}
	}

	for input > 0 {
		switch {
		case input >= 1000:
			roman_numeral += "M"
			input -= 1000

		case input >= 500:
			calcNumber(900, 500)

		case input >= 100:
			calcNumber(400, 100)

		case input >= 50:
			calcNumber(90, 50)

		case input >= 10:
			calcNumber(40, 10)

		case input >= 5:
			calcNumber(9, 5)

		case input >= 1:
			calcNumber(4, 1)
		}
	}

	return roman_numeral, nil
}
