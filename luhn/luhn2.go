package luhn

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Valid2(id string) bool {
	id = strings.Replace(id, " ", "", -1)

	length := utf8.RuneCount([]byte(id))

	if length <= 1 {
		return false
	}

	sum := 0
	for i, n := range id {

		if !unicode.IsDigit(n) {
			return false
		}

		v := int(n - '0')

		if (i%2 == 0 && length%2 == 0) || (i%2 != 0 && length%2 != 0) {
			v *= 2
			if v > 9 {
				v -= 9
			}

		}

		sum += v
	}
	return sum%10 == 0
}
