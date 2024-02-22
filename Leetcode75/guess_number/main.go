package main

import "math/rand"

func main() {
}

func guessNumber(n int) int {
	low, high := 1, n

	for low <= high {
		mid := low + (high-low)/2

		result := guess(mid, n)

		if result == 0 {
			return mid
		} else if result > 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func guess(k, n int) int {
	pick := rand.Intn(n)

	if pick == k {
		return 0
	} else if pick > k {
		return 1
	} else {
		return -1
	}
}
