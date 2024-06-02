package cryptosquare

import (
	"strings"
	"unicode"
)

func Encode(pt string) string {
	normalized := strings.Map(
		func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return r
			}
			return -1
		}, pt)

	normalized = strings.ToLower(normalized)

	if len(normalized) == 0 {
		return ""
	}

	r, c := 1, 1

	for r*c < len(normalized) {
		if r == c {
			c++
		} else {
			r++
		}
	}

	encoded := ""

	for i := 0; i < c; i++ {
		for j := 0; j < r*c; j += c {
			if j+i < len(normalized) {
				encoded += string(normalized[i+j])
			} else {
				encoded += " "
			}
		}
		if i < c-1 {
			encoded += " "
		}
	}

	return encoded
}
