package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var regex *regexp.Regexp = regexp.MustCompile(`\w+'?\w+|\w{1}`)

func WordCount(phrase string) Frequency {

	words := regex.FindAllString(strings.ToLower(phrase), -1)

	freq := Frequency{}

	for _, w := range words {
		freq[w]++
	}

	return freq
}
