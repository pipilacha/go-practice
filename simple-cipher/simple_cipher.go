package cipher

import (
	"strings"
)

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

// NewCaesar creates and returns a shift cipher with distance  of 3
func NewCaesar() Cipher {
	return shift{distance: 3}
}

// NewCaesar creates a shift cipher with a arg supplied distance
//
//	the permitted value for distance are values in ranges [-25,-1] or [1,25]
//	returns nil if distance is out of range
func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}

	return shift{distance: distance}
}

// shiftRight shift the specied rune to the right, it wraps around
func shiftRight(dist int, r rune) (shifted rune) {
	shifted = r + rune(dist)
	if shifted > 'z' {
		shifted -= 26
	}
	return
}

// shiftRight shift the specied rune to the left, it wraps around
func shiftLeft(dist int, r rune) (shifted rune) {
	shifted = r - rune(dist)
	if shifted < 'a' {
		shifted += 26
	}
	return
}

// Encode receives a raw input (string) and encodes it using the distance set and returns it.
//
//	By default encode shifts the rune to the right, but if the shift distance is negative it will shift to the left
func (c shift) Encode(input string) (encoded string) {
	direction := shiftRight

	if c.distance < 0 {
		direction = shiftLeft
		c.distance *= -1
	}

	for _, r := range strings.ToLower(input) {

		if r >= 'a' && r <= 'z' {
			encoded += string(direction(c.distance, r))
		}
	}

	return encoded
}

// Decode receives a encoded input (string) and decodes it using the distance set and returns it.
//
//	By default decode shifts the rune to the left, but if the shift distance is negative it will shift to the right
func (c shift) Decode(input string) (decoded string) {
	direction := shiftLeft

	if c.distance < 0 {
		direction = shiftRight
		c.distance *= -1
	}

	for _, r := range input {
		decoded += string(direction(c.distance, r))
	}

	return decoded
}

// NewVigenere creates a new cipher with the specied key
//
//	returns nil if key is empty, contains only a or contains values not in [a-z]
func NewVigenere(key string) Cipher {
	number_of_a := 0

	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		if r == 'a' {
			number_of_a++
		}
	}

	if number_of_a == len(key) {
		return nil
	}

	return vigenere{key: key}
}

// Encode receives a raw input (string) and encodes it using the set key and returns it.
func (v vigenere) Encode(input string) (encoded string) {
	i := 0
	for _, r := range strings.ToLower(input) {
		if r >= 'a' && r <= 'z' {
			encoded += string(shiftRight(int(v.key[i]-'a'), r))
			i++
			if i == len(v.key) {
				i = 0
			}
		}
	}
	return encoded
}

// Decode receives a encoded input (string) and decodes it using the set key and returns it.
func (v vigenere) Decode(input string) (decoded string) {
	i := 0
	for _, r := range input {
		decoded += string(shiftLeft(int(v.key[i]-'a'), r))
		i++
		if i == len(v.key) {
			i = 0
		}
	}
	return decoded
}
