package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	return
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() (dice int) {
	rand.Seed(time.Now().UnixNano())
	dice = rand.Intn(21-1) + 1 // define lower and upper bound
	return
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() (wand float64) {
	rand.Seed(time.Now().UnixNano())
	wand = 0.0 + rand.Float64()*(12.0-0.0) //define upper and lower bound
	return

}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() (a []string) {
	a = []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Seed(time.Now().UnixNano())
	// Shuffling slice or arrays
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return

}

func main() {
	SeedWithTime()
	fmt.Println(RollADie())
	fmt.Println(GenerateWandEnergy())
	fmt.Println(ShuffleAnimals())
}
