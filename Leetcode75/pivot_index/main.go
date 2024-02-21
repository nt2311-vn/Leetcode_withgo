package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{2, -1, 1}))
}

func pivotIndex(nums []int) int {
	/*
		for i := 0; i < len(nums); i++ {
			leftVal := 0
			for _, v := range nums[0:i] {
				leftVal += v
			}

			rightVal := 0
			for _, v := range nums[i+1:] {
				rightVal += v
			}

			if leftVal == rightVal {
				return i
			}
		}
		return -1
	*/
	// Complexity of O(n)
	totalSum := 0

	for _, v := range nums {
		totalSum += v
	}
	leftSum := 0
	for index, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return index
		}
		leftSum += num
	}
	return -1
}
