package summultiples

func SumMultiples(limit int, divisors ...int) (total int) {
	multipliers := map[int]bool{}

	for _, item := range divisors {
		for i := item; i < limit; i++ {
			if i == 0 {
				break
			}

			if i > 0 && i%item == 0 {
				multipliers[i] = true
			}
		}
	}

	for k := range multipliers {
		total += k
	}

	return
}
