package main

import "fmt"

func main() {
	fmt.Printf("Expected value: 3, actual: %d\n", longestSubArray([]int{1, 1, 0, 1}))
	fmt.Printf("Expected value: 5, actual: %d\n", longestSubArray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Printf("Expected value: 2, actual: %d\n", longestSubArray([]int{1, 1, 1}))
}

func longestSubArray(nums []int) int {
	var maxLen, zeroCount, start int

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}

		currentLen := end - start

		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	if maxLen == len(nums) {
		return maxLen - 1
	}
	return maxLen
}
