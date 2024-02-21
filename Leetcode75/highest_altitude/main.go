package main

import "fmt"

func main() {
	fmt.Println(largetstAltitude([]int{-5, 1, 5, 0, -7}))
}

func largetstAltitude(gain []int) int {
	maxAlt := 0

	// Assing initial value to max to find
	if gain[0] > 0 {
		maxAlt = gain[0]
	}

	currentAlt := 0
	for i := 0; i < len(gain); i++ {
		currentAlt += gain[i]
		if maxAlt < currentAlt {
			maxAlt = currentAlt
		}
	}
	return maxAlt
}
