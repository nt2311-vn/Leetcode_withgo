package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Expected value: true, actual: %v\n", close_string("abc", "bca"))
	fmt.Printf("Expected value: false, actual: %v\n", close_string("a", "aa"))
	fmt.Printf("Expected value: true, actual: %v\n", close_string("cabbba", "abbccc"))
}

func close_string(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charCount1 := make(map[rune]int)
	charCount2 := make(map[rune]int)

	for _, char := range s1 {
		charCount1[char]++
	}

	for _, char := range s2 {
		charCount2[char]++
	}

	if len(charCount1) != len(charCount2) {
		return false
	}

	for char := range charCount1 {
		if _, exist := charCount2[char]; !exist {
			return false
		}
	}

	freq1 := make([]int, 0, len(charCount1))
	freq2 := make([]int, 0, len(charCount2))

	for _, count := range charCount1 {
		freq1 = append(freq1, count)
	}

	for _, count := range charCount2 {
		freq2 = append(freq2, count)
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := range freq1 {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
