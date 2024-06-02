package pangram

import (
	"slices"
	"strings"
)

func IsPangram(input string) bool {
	alphabet := make([]bool, 26)
	for _, r := range strings.ToLower(input) {
		switch r {
		case 'a':
			alphabet[0] = true
		case 'b':
			alphabet[1] = true
		case 'c':
			alphabet[2] = true
		case 'd':
			alphabet[3] = true
		case 'e':
			alphabet[4] = true
		case 'f':
			alphabet[5] = true
		case 'g':
			alphabet[6] = true
		case 'h':
			alphabet[7] = true
		case 'i':
			alphabet[8] = true
		case 'j':
			alphabet[9] = true
		case 'k':
			alphabet[10] = true
		case 'l':
			alphabet[11] = true
		case 'm':
			alphabet[12] = true
		case 'n':
			alphabet[13] = true
		case 'o':
			alphabet[14] = true
		case 'p':
			alphabet[15] = true
		case 'q':
			alphabet[16] = true
		case 'r':
			alphabet[17] = true
		case 's':
			alphabet[18] = true
		case 't':
			alphabet[19] = true
		case 'u':
			alphabet[20] = true
		case 'v':
			alphabet[21] = true
		case 'w':
			alphabet[22] = true
		case 'x':
			alphabet[23] = true
		case 'y':
			alphabet[24] = true
		case 'z':
			alphabet[25] = true
		}
	}
	return len(slices.Compact(alphabet)) == 1 && alphabet[0]
}
