package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(height))
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	left, right, currentMax := 0, len(height)-1, 0

	for left < right {
		if height[left] < height[right] {
			currentMax = maxNum(currentMax, (right-left)*height[left])
			left++
		} else {
			currentMax = maxNum(currentMax, (right-left)*height[right])
			right--
		}
	}

	return currentMax
}
