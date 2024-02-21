package main

import "fmt"

func main() {
	fmt.Printf("Expected value: 2, actutal: %d\n", maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Printf("Expected value: 1, actual: %d\n", maxOperations([]int{3, 1, 3, 4, 3}, 6))
}

func maxOperations(nums []int, k int) int {
	countMap := make(map[int]int)
	countOps := 0

	for _, num := range nums {
		complement := k - num

		if countMap[complement] > 0 {
			countOps++
			countMap[complement]--
		} else {
			countMap[num]++
		}
	}

	return countOps
}
