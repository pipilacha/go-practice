package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	normalized := 0
	for _, digit := range inputDigits {

		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")

		}
		// 2365
		// 0*10+2 = 2
		// 2*10+3 = 23
		// 23*10+6 = 236
		// 236*10+5 = 2365
		normalized = normalized*inputBase + digit
	}

	if normalized == 0 {
		return []int{0}, nil
	}

	outputDigits := []int{}
	for normalized > 0 {
		outputDigits = append([]int{normalized % outputBase}, outputDigits...)
		normalized = normalized / outputBase
	}

	return outputDigits, nil
}
