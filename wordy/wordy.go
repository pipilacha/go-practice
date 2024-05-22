package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (result int, ok bool) {
	operation_re := regexp.MustCompile(`((?:-?\d+ (?:divided by|multiplied by|plus|minus) )+-?\d+\?$)|What is (-?\d+\?)`)
	operation := operation_re.FindStringSubmatch(question)

	if len(operation) == 0 {
		return 0, false
	}

	pos := 1
	if operation[2] != "" {
		pos = 2
	}

	list := strings.Split(strings.TrimSuffix(strings.Replace(operation[pos], "d b", "d-b", -1), "?"), " ")

	result, _ = strconv.Atoi(list[0])

	for i := 1; i < len(list); i += 2 {
		x, _ := strconv.Atoi(list[i+1])
		switch list[i] {
		case "plus":
			result += x
		case "minus":
			result -= x
		case "divided-by":
			result /= x
		case "multiplied-by":
			result *= x
		}
	}

	return result, true
}
