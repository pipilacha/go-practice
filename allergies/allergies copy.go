package allergies

import (
	"fmt"
)

// // The scoring system follows a binary format, the first test are eggs so:
// egg 		 = 0, 2 ^ 0 = 1
// peanuts 	 = 1, 2 ^ 1 = 2
// shellfish 	 = 2, 2 ^ 2 = 4
// strawberries = 3, 2 ^ 3 = 8
// and so on...
// at the end the scores are just added.
// By converting the score back to binary we can just iterate over it and see if that person is allergic
// to that allergen

var allergiesCodes2 = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies2(allergies uint) (allrgs []string) {
	if allergies == 0 {
		return
	}

	binAllergies := fmt.Sprintf("%08b", allergies) // transform the score to binary

	for i := 0; i <= 7; i++ { // read from right to left, we are only concerned about the last 7
		if binAllergies[len(binAllergies)-1-i] == '1' {
			allrgs = append(allrgs, allergiesCodes2[i])
		}
	}
	return
}

func AllergicTo2(allergies uint, allergen string) bool {
	if allergies == 0 && allergen != "" {
		return false
	}

	binAllergies := fmt.Sprintf("%08b", allergies)

	index := 0
	for i, v := range allergiesCodes2 {
		if v == allergen {
			index = i
			break
		}
	}

	return binAllergies[len(binAllergies)-1-index] == '1'
}

/*
BenchmarkAllergies-8              107479             10182 ns/op            3393 B/op        108 allocs/op
BenchmarkAllergicTo-8             276958              4268 ns/op             256 B/op         32 allocs/op
PASS
ok      allergies       2.441s
*/
