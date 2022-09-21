package main

import (
	"fmt"
	"math/rand"
)

// create a random.NextBool() function
// that returns true or false with equal probability

func NextBool() bool {
	return rand.Intn(2) == 1
}

func main() {
	var size, seed int
	fmt.Scanln(&size, &seed)

	rand.Seed(int64(seed))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if NextBool() {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
