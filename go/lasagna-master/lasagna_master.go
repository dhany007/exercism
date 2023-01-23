package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, average int) int {
	if average == 0 {
		average = 2
	}

	return len(layers) * average
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	var (
		grNoodles   int     = 50
		nameNoodles string  = "noodles"
		grSauce     float64 = 0.2
		nameSauce   string  = "sauce"
	)

	for i := 0; i < len(layers); i++ {
		if layers[i] == nameNoodles {
			noodles += grNoodles
		}

		if layers[i] == nameSauce {
			sauce += grSauce
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipes []float64, portions int) (result []float64) {
	for i := 0; i < len(recipes); i++ {
		value := recipes[i] * (float64(portions) / 2)
		result = append(result, value)
	}

	return
}

