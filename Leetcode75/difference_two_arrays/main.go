package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}

func findDifference(nums1, nums2 []int) [][]int {
	result := [][]int{}
	distArr1 := []int{}
	distArr2 := []int{}

	for _, v1 := range nums1 {
		if !checkItemExist(nums2, v1) && !checkItemExist(distArr1, v1) {
			distArr1 = append(distArr1, v1)
		}
	}

	for _, v2 := range nums2 {
		if !checkItemExist(nums1, v2) && !checkItemExist(distArr2, v2) {
			distArr2 = append(distArr2, v2)
		}
	}

	result = append(result, distArr1)
	result = append(result, distArr2)
	return result
}

func checkItemExist(intArr []int, item int) bool {
	for _, v := range intArr {
		if v == item {
			return true
		}
	}
	return false
}
