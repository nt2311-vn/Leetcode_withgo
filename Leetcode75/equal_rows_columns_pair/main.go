package main

import (
	"fmt"
)

func main() {
	fmt.Printf(
		"Expected value: 1, actual: %d\n",
		equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}),
	)
	fmt.Printf(
		"Expected value: 3, actual: %d\n",
		equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}),
	)
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	count := 0

	rowStrs := make([]string, n)
	for i := 0; i < n; i++ {
		rowStr := ""
		for j := 0; j < n; j++ {
			rowStr += fmt.Sprintf("%d,", grid[i][j])
		}
		rowStrs[i] = rowStr
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			colStr := ""
			for k := 0; k < n; k++ {
				colStr += fmt.Sprintf("%d,", grid[k][j])
			}
			if colStr == rowStrs[i] {
				count++
			}
		}
	}

	return count
}
