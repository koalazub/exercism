package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (n int, s float64) {

	for _, k := range layers {
		if k == "noodles" {
			n += 50
		} else if k == "sauce" {
			s += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend_ingredients, ingredients []string) {
	last_item := friend_ingredients[len(friend_ingredients)-1]
	ingredients[len(ingredients)-1] = last_item
}

// Scale takes in two types of portions - one is the original portion, the parameter is the
// portion you want to make for. The OG portion starts at 2
func ScaleRecipe(quantities []float64, portions int) []float64 {
	sum_portions := make([]float64, len(quantities))
	for i := range quantities {
		sum_portions[i] = quantities[i] * (float64(portions) / 2.0)
	}

	return sum_portions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
