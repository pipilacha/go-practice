// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) (acronym string) {
	ignore_next := false // ignore next will tell the function to match next letter after "'", split at there

	splitted := strings.FieldsFunc(s, func(r rune) bool {
		if ignore_next {
			ignore_next = false
			return true
		}

		if r == '\'' {
			ignore_next = true
		}
		return r < 'A' || r > 'Z' && r < 'a' || r > 'z'
	})

	for _, v := range splitted {
		acronym += strings.ToUpper(string(v[0]))
	}

	return
}

func Abbreviate2(s string) (out string) {

	s = strings.Replace(s, "-", " ", -1)

	s = strings.Replace(s, "_", " ", -1)

	words := strings.Fields(s)

	for i := range words {

		out += string(words[i][0])

	}

	return strings.ToUpper(out)

}

/*
go test -benchmem -run=^$ -bench ^BenchmarkAcronym* acronym
cpu: AMD Ryzen 5 7520U with Radeon Graphics
BenchmarkAcronym-8        239586              4489 ns/op            1200 B/op        101 allocs/op
BenchmarkAcronym2-8       267073              4966 ns/op            1200 B/op        101 allocs/op
PASS
ok      acronym 3.47
*/
