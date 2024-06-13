package anagram

import (
	"fmt"
	"strings"
)

func getWordByteSum2(word string) (total rune, biggest rune, shifts int) {
	for _, r := range word {
		total += r

		shifts += 1 << (r / 2)

		if r > biggest {
			biggest = r
		}
	}
	fmt.Println(word, total, biggest, shifts)
	return
}

func Detect2(subject string, candidates []string) (selected []string) {

	subject = strings.ToLower(subject)

	subject_sum, subject_biggest, subject_shifts := getWordByteSum2(subject)

	for _, word := range candidates {
		lowered := strings.ToLower(word)

		word_sum, word_biggest, word_shifts := getWordByteSum2(lowered)

		if lowered != subject && word_sum == subject_sum && word_biggest == subject_biggest && word_shifts == subject_shifts {
			selected = append(selected, word)
		}
	}

	return
}
