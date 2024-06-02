package grains

import "fmt"

var bypass = false

func Square2(number int) (uint64, error) {
	if number < 1 || number > 64 && !bypass {
		return 0, fmt.Errorf("invalid square number: %v. only 1-64 allowed", number)
	}

	bypass = false

	bin := 0b01 << (number - 1)

	return uint64(bin), nil
}

func Total2() uint64 {
	bypass = true
	t, _ := Square(65)
	t -= 1
	return t
}
