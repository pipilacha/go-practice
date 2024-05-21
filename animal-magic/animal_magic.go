package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	// for {
	// 	val := rand.Intn(21)
	// 	if val > 0 {
	// 		return val
	// 	}
	// }
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	// for {
	// 	val := rand.Float64()
	// 	if val < 0.12 {
	// 		return val * 100
	// 	}
	// }
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(8, func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}
