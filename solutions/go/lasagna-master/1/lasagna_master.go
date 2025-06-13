package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
    switch minutes {
        case 0:
    		return len(layers) * 2
        default:
    		return len(layers) * minutes
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, v := range layers {
        switch v {
            case "noodles":
        		noodles += 50
            case "sauce":
        		sauce += 0.2
        }
    }
	return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	var lastIndexFriendsList int = len(friendsList) - 1
    var lastIndexMyList int = len(myList) - 1

    myList[lastIndexMyList] = friendsList[lastIndexFriendsList]
    
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
    var size = len(amounts)    
    var newAmounts = make([]float64, size)
    var koef float64 = float64(portions) / 2
    copy(newAmounts, amounts)
	for i, _ := range newAmounts {
        newAmounts[i] *= koef
    }
    return newAmounts
}
