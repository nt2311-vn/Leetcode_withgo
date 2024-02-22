package main

import "fmt"

func main() {
	fmt.Printf(
		"Expected value: 6, actual: %d\n",
		longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2),
	)

	fmt.Printf(
		"Expected value: 10, actual: %d\n",
		longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3),
	)
}

func longestOnes(nums []int, k int) int {
	currentMax := 0
	currentCount := 0

	for _, val := range nums {
		if val == 1 {
			currentCount++
		} else {
			currentMax = maxNum(currentMax, currentCount)
			currentCount++
			k--
		}
	}

	return currentMax
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
