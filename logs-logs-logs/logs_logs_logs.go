package logs

import (
	"fmt"
	"unicode/utf8"
)

var cp_map map[string]string = map[string]string{
	"U+2757":  "recommendation",
	"U+1F50D": "search",
	"U+2600":  "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		v, ok := cp_map[fmt.Sprintf("%U", r)]
		if ok {
			return v
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) (new_string string) {
	for _, r := range log {
		if r == oldRune {
			new_string += fmt.Sprintf("%c", newRune)
		} else {
			new_string += fmt.Sprintf("%c", r)
		}
	}
	return new_string
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
