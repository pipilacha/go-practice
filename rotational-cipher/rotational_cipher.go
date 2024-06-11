package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	cipher := ""

	for _, v := range plain {
		if !unicode.IsLetter(v) {
			cipher += string(v)
			continue
		}

		shift := int(v) + shiftKey

		if unicode.IsUpper(v) && shift > 'Z' || unicode.IsLower(v) && shift > 'z' {
			shift -= 26
		}

		cipher += string(rune(shift))
	}

	return cipher
}
