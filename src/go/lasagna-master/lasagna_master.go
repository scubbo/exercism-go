package lasagna

func PreparationTime(layers []string, prep_time_per_layer int) int {
	if prep_time_per_layer == 0 {
		prep_time_per_layer = 2
	}
	return len(layers) * prep_time_per_layer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(friend_list []string, own_recipe []string) {
	own_recipe[len(own_recipe)-1] = friend_list[len(friend_list)-1]
}

// The above passes the tests, but I don't know why it does.
// I thought the whole point of pointers (as introduced in this
// section) was that they were the only way to have a function make
// a change outside its own scope? Shouldn't the answer be:
//
//func AddSecretIngredient(friend_list []string, own_recipe *[]string) {
//	// https://flaviocopes.com/golang-does-not-support-indexing/
//	(*own_recipe)[len(*own_recipe)-1] = friend_list[len(friend_list)-1]
//}

func ScaleRecipe(amounts_to_cook_two_portions []float64, num_portions int) []float64 {
	var output []float64
	if len(amounts_to_cook_two_portions) == 0 {
		return output
	}
	for i := 0; i < len(amounts_to_cook_two_portions); i++ {
		// I miss Python...
		output = append(output, amounts_to_cook_two_portions[i]*float64(num_portions)/2)
	}
	return output
}
