package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	mapNum := map[int]int{}

	for _, val := range nums {
		mapNum[val]++

		if mapNum[val] > 1 {
			delete(mapNum, val)
		}
	}
	singleNumber := 0

	for key := range mapNum {
		singleNumber = key
	}
	return singleNumber
}
