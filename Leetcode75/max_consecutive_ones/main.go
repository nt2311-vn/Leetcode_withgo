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
	left, maxLen, zeroCount := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		currentWindowLen := right - left + 1
		if currentWindowLen > maxLen {
			maxLen = currentWindowLen
		}
	}
	return maxLen
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
