package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) (song []string) {
	n_text := map[int]string{
		0:  "No",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
	}

	bottles_or_bottle := func(n int) string {
		if n == 1 {
			return "bottle"
		}
		return "bottles"
	}

	for i := startBottles; i > startBottles-takeDown; i-- {
		song = append(song,
			fmt.Sprintf("%s green %s hanging on the wall,", n_text[i], bottles_or_bottle(i)),
			fmt.Sprintf("%s green %s hanging on the wall,", n_text[i], bottles_or_bottle(i)),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(n_text[i-1]), bottles_or_bottle(i-1)),
		)

		if i-1 > startBottles-takeDown {
			song = append(song, "")
		}
	}

	return song
}
