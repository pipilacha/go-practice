package scrabble

import (
	"strings"
)

func getLetterValue(letter string) int {
	switch letter {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	default: // if letter = "Q" or "Z"
		return 10
	}
}

func Score(word string) (total int) {
	for _, v := range word {
		total += getLetterValue(strings.ToUpper(string(v)))
	}
	return
}
