package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	max1, max2 := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= max1 {
			max1 = num
		} else if num <= max2 {
			max2 = num
		} else {
			return true
		}
	}

	return false
}
