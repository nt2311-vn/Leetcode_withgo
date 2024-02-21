package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits(5))
}

func countBits(n int) []int {
	res := make([]int, n+1)
	t := 1
	for i := 1; i <= n; i++ {
		if i == t<<1 {
			t = t << 1
		}
		res[i] = 1 + res[i-t]
	}

	return res
}
