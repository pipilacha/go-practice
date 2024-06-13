package prime

func Factors(n int64) (factors []int64) {
	factor := int64(2)

	for n > 1 {
		if n%factor == 0 {
			factors = append(factors, factor)
			n /= factor
		} else {
			factor++
		}
	}

	return
}
