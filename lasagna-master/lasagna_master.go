package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}
	return len(layers) * avg
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friends []string, mine []string) {
	mine[len(mine)-1] = friends[len(friends)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(layers []float64, portions int) []float64 {
	new_layers := make([]float64, len(layers))
	copy(new_layers, layers)
	ratio := float64(portions) / 2
	for i, v := range new_layers {
		new_layers[i] = v * ratio
	}
	return new_layers
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
