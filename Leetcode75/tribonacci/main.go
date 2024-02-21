package main

import "fmt"

func main() {
	fmt.Println(tribonaci(25))
}

func tribonaci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	a, b, c := 0, 1, 1

	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
