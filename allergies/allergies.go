package allergies

var allergiesCodes = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(allergies uint) (allergs []string) {

	for i := 0; i < 8 && allergies > 0; i++ { // we will shift allergies 1 bit to the right until alergies > 0. we start from 0, bit 0
		if allergies&1 == 1 { // by using Bitwise AND it can tell us if that bit (0,1,2,4,8, etc) is on
			allergs = append(allergs, allergiesCodes[i])
		}
		allergies = allergies >> 1 // then we shift allergies 1 bit to the right

	}

	return
}

func AllergicTo(allergies uint, allergen string) bool {
	if allergies == 0 {
		return false
	}

	for i, v := range allergiesCodes { // we will iterate over the allergies code array
		if v == allergen {
			return allergies&(1<<i) != 0 // we shift left 1 n bits accordingly to the positiong of the allergen in the allergiesCodes, eg. 1 << 0 if alergen == eggs
			// bitwise returns the same bits turned on between the two,
			// and since we are passing 1,2,4,8,16,32,...
			// allergies will either contain the entire bits of of the allergen or will have none in common (will return 0)
			// eg: 34 is 00100010, 32 is 00100000, 16 is 00010000, 00000010
			//	34&32			34&16			34&2
			//	00100010		00100010		00100010
			//	00100000 =		00010000 =		00000010 =
			//	00100000(32)	00000000(0)		00000010(2)
		}
	}

	return false
}

/*
BenchmarkAllergies-8              293902              4058 ns/op            3136 B/op         76 allocs/op
BenchmarkAllergicTo-8            6519948               188.2 ns/op             0 B/op          0 allocs/op
PASS
*/
