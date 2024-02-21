package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurence([]int{3, 5, -2, -3, -6, -6}))
}

func uniqueOccurence(arr []int) bool {
	storeOccur := map[int]int{}

	for _, v := range arr {
		storeOccur[v]++
	}

	return checOccur(storeOccur)
}

func checOccur(mapNum map[int]int) bool {
	checkOccur := map[int]bool{}

	for _, value := range mapNum {
		if checkOccur[value] {
			return false
		}

		checkOccur[value] = true
	}

	return true
}
