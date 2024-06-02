package grains

import (
	"fmt"
)

func Square3(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("invalid square number: %v. only 1-64 allowed", number)
	}

	bin := 0b01 << (number - 1)

	return uint64(bin), nil
}

func Total3() uint64 {
	t := uint64(0)
	for i := 1; i < 65; i++ {
		v, _ := Square(i)
		t += v
	}
	return t
}
