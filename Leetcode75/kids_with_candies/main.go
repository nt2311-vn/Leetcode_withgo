package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}

	findMax := func(number []int) int {
		maxVal := 0

		for _, val := range number {
			if val > maxVal {
				maxVal = val
			}
		}
		return maxVal
	}(candies)

	for _, val := range candies {
		if val+extraCandies >= findMax {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
