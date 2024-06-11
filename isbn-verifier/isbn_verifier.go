package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {

	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	number := 0

	for i, v := range isbn {

		val := 0

		if v != 'X' {
			parsed, err := strconv.Atoi(string(v))
			if err != nil {
				return false
			}
			val = parsed

		} else {
			if i != 9 {
				return false
			}
			val = 10
		}

		number += val * (10 - i)
	}

	return number%11 == 0
}
