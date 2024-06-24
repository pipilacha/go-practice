package encode

import (
	"strconv"
	"strings"
)

func encodeAmount(c int, r rune) (encoded string) { // transform AAA to 3A or A to A
	if c > 1 {
		encoded += strconv.Itoa(c) + string(r)
	} else {
		encoded += string(r)
	}
	return
}

func RunLengthEncode(input string) (encoded string) {
	if len(input) == 0 {
		return
	}

	counter := 1
	prev := rune(input[0])
	for i, r := range input {
		if i == 0 {
			continue
		}

		if i == len(input)-1 { // when we are at the last letter we need to check if it is the same as the previous
			if prev == r {
				counter++
				encoded += encodeAmount(counter, r)
				continue
			}
			encoded += encodeAmount(counter, prev)
			encoded += encodeAmount(1, r)
			continue
		}

		if prev == r { // is the current letter == previous we increase counter an continue
			counter++
			continue
		}

		encoded += encodeAmount(counter, prev) // when letter is not the same we encode the previous with the number of occurrences
		counter = 1
		prev = r

	}
	return encoded
}

func RunLengthDecode(input string) (decoded string) {
	if len(input) == 0 {
		return ""
	}

	index := 0
	amount := "0"
	for index < len(input) {
		isNumber := '1' <= input[index] && '9' >= input[index] // we check if current run is a number

		if isNumber { // if it is a number we increase we add it to the amounr of the next letter, increase index and continue
			amount += string(input[index]) // if encoded = 123A , amount = "0123" and we use that amount to repeat the letter
			index++
			continue
		}

		// if it is not a number we decode
		if amount == "0" { // if amount = 0 this means that letter only has 1 occurence, so we set it to one
			amount = "1"
		}
		n, _ := strconv.Atoi(amount)
		decoded += strings.Repeat(string(input[index]), n)
		amount = "0"
		index++
	}
	return
}

/*
BenchmarkRunLengthEncode-8        734004              1617 ns/op             224 B/op         41 allocs/op
BenchmarkRunLengthDecode-8        367881              2749 ns/op             632 B/op         89 allocs/op
PASS
ok      encode  2.256s
*/
