package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	if utf8.RuneCount([]byte(a)) != utf8.RuneCount([]byte(b)) {
		return 0, errors.New("strands must be of the same size")
	}

	count := 0

	for i, v := range a {
		if b[i] != byte(v) {
			count++
		}
	}
	return count, nil
}
