package main

import "fmt"

func main() {
	moveZeroes([]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1})
}

func moveZeroes(nums []int) {
	pos := 0
	for _, val := range nums {
		if val != 0 {
			nums[pos] = val
			pos++
		}
	}

	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}
