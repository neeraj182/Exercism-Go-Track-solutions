
package main

import (
	"fmt"
)

func PreparationTime(layers []string, time int) int {
	var total_time int
	total_time = len(layers) * time

	if time == 0 {
		total_time = len(layers) * 2
		return total_time
	}
	return total_time
}

// one way to perform is tp use strings.contains(), but need to include strings package
/*func Quantities(layers []string) (total_noodle int, total_sauce float64) {
noodleqt := 50
sauceqt := 0.2
for _, layer := range layers {
	if strings.Contains(layer, "noodle") {
		total_noodle += noodleqt
	} else if strings.Contains(layer, "sauce") {
		total_sauce += sauceqt
	}
}
return
*/

// the other way
func Quantities(layers []string) (total_noodle int, total_sauce float64) {
	noodleqt := 50
	sauceqt := 0.2
	for _, layer := range layers {
		switch layer {
		case "noodles":
			total_noodle += noodleqt
		case "sauce":
			total_sauce += sauceqt
		}
	}
	return

}

func AddSecretIngredient(friendList []string, myList []string) []string {
	lastElement := friendList[len(friendList)-1] //index of the last element in slice
	myList[len(myList)-1] = lastElement
	return myList

}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newQuantities []float64
	for _, quantity := range quantities {
		newQuantities = append(newQuantities, quantity/2*float64(portions))
	}
	return newQuantities

}

func main() {
	layers := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	quantities := []float64{1.2, 3.6, 10.5}

	fmt.Println(PreparationTime(layers, 0))
	fmt.Println(Quantities(layers))
	fmt.Println(AddSecretIngredient(friendsList, layers))
	fmt.Println(ScaleRecipe(quantities, 4))

}
