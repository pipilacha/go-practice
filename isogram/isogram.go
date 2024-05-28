package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	letters := map[rune]bool{}
	for _, c := range strings.ToLower(word) {
		if letters[c] && unicode.IsLetter(c) {
			return false
		}
		letters[c] = true
	}
	return true
}
