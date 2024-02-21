package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	currentSum := 0

	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	for i := k; i < len(nums); i++ {
		currentSum += nums[i] - nums[i-k]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return float64(maxSum) / float64(k)
}
