package lasagna

func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodleCount := 0
	sauceCount := 0
	for _, l := range layers {
		switch l {
		case "noodles":
			noodleCount++

		case "sauce":
			sauceCount++
		}
	}

	return noodleCount * 50, float64(sauceCount) * 0.2
}

func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaledAmounts := make([]float64, len(quantities))

	for i, quantity := range quantities {
		scaledAmounts[i] = quantity / 2 * float64(numPortions)
	}

	return scaledAmounts
}
