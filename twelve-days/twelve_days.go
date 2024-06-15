package twelve

import "fmt"

func Verse(i int) string {
	day := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}

	gaveme := []string{
		"a Partridge in a Pear Tree.",
		"two Turtle Doves, and ",
		"three French Hens, ",
		"four Calling Birds, ",
		"five Gold Rings, ",
		"six Geese-a-Laying, ",
		"seven Swans-a-Swimming, ",
		"eight Maids-a-Milking, ",
		"nine Ladies Dancing, ",
		"ten Lords-a-Leaping, ",
		"eleven Pipers Piping, ",
		"twelve Drummers Drumming, ",
	}

	got := ""
	for j := 0; j < i; j++ {
		got = gaveme[j] + got
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", day[i-1], got)
}

func Song() (song string) {
	for i := 1; i < 13; i++ {
		song += Verse(i)

		if i < 12 {
			song += "\n"
		}
	}
	return
}
