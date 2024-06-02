package grains

import (
	"fmt"
)

func Square4(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("invalid square number: %v. only 1-64 allowed", number)
	}

	bin := 0b01 << (number - 1)

	return uint64(bin), nil
}

func Total4() uint64 {
	t, _ := Square(64)
	return t*2 - 1
}
