package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span <= 0 || span > len(digits) {
		return 0, errors.New("span cannot be lower than one or greater than the length of the string")
	}

	highest := 0

	for i := 0; i <= len(digits)-span; i++ {
		product := 1

		for _, digit := range digits[i : i+span] {
			value, notdigit_error := strconv.Atoi(string(digit))

			if notdigit_error != nil {
				return 0, errors.New("digits input must only contain digits")
			}

			product *= value
		}

		if product > highest {
			highest = product
		}
	}

	return int64(highest), nil
}
