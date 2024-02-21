package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Expect max vowels is 3, actual: %d\n", maxVowels("abciiidef", 3))
	fmt.Printf("Expect max vowels is 2, acutal: %d\n", maxVowels("aiueo", 2))
	fmt.Printf("Expect max vowels is 2, actual: %d\n", maxVowels("leetcode", 3))
	fmt.Printf("Expect max vowels is 4, actual: %d\n", maxVowels("weallloveyou", 7))
}

func maxVowels(s string, k int) int {
	vowels := map[rune]bool{'a': true, 'i': true, 'o': true, 'e': true, 'u': true}
	currentMax := 0
	currentCount := 0

	for i, ch := range s {
		if vowels[ch] {
			currentCount++
		}

		if i >= k {
			if vowels[rune(s[i-k])] {
				currentCount--
			}
		}
		if currentCount > currentMax {
			currentMax = currentCount
		}
	}
	return currentMax
}
