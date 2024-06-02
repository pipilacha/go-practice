// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	amount_letters, all_capitalized := 0, 0
	var last rune
	for _, r := range remark {
		if unicode.IsLetter(r) {
			amount_letters++
			if r >= 'A' && r <= 'Z' {
				all_capitalized++
			}
		}
		last = r
	}

	if amount_letters == 0 {
		amount_letters = -1
	}

	switch {
	case last == '?' && amount_letters != all_capitalized:
		return "Sure."
	case last == '?' && amount_letters == all_capitalized:
		return "Calm down, I know what I'm doing!"
	case amount_letters == all_capitalized:
		return "Whoa, chill out!"
	default:
		return "Whatever."

	}
}
