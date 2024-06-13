package atbash

import (
	"strings"
)

func Atbash(s string) (cipher string) {
	s = strings.ToLower(s)

	added := 0

	for _, v := range s {

		if v >= 'a' && v <= 'z' || v >= '0' && v <= '9' {
			if added == 5 {
				cipher += " "
				added = 0
			}

			if v >= 'a' {
				cipher += string('z' - (v - 'a'))
			} else {
				cipher += string(v)
			}

			added++
		}
	}
	return
}
