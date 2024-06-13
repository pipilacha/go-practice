package anagram

import (
	"sort"
	"strings"
)

func sortWord(word string) string {
	split := strings.Split(word, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func Detect(subject string, candidates []string) (selected []string) {

	subject = strings.ToLower(subject)

	subject_sorted := sortWord(subject)

	for _, word := range candidates {
		candidate := strings.ToLower(word)

		word_sorted := sortWord(candidate)

		if candidate != subject && word_sorted == subject_sorted {
			selected = append(selected, word)
		}
	}

	return
}
