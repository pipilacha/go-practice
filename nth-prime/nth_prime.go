package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("there is no zeroth prime")
	}

	primes := []int{2, 3, 5, 7, 11, 13}

	if n <= 6 {
		return primes[n-1], nil
	}

	nth := 6
	next := 13

	for nth < n {
		next++

		next_is_prime := true

		for _, v := range primes[:nth] {
			if next%v == 0 {
				next_is_prime = false
				break
			}
		}

		if next_is_prime {
			primes = append(primes, next)
			nth++
		}
	}
	return next, nil
}
