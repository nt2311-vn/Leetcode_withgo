package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}
	vowels := "aiueo"
	breakStr := strings.Split(s, "")

	i, j := 0, len(breakStr)-1
	for i < j {
		if !strings.Contains(vowels, strings.ToLower(breakStr[i])) {
			i++
			continue
		}
		if !strings.Contains(vowels, strings.ToLower(breakStr[j])) {
			j--
			continue
		}

		breakStr[i], breakStr[j] = breakStr[j], breakStr[i]
		i++
		j--
	}

	return strings.Join(breakStr, "")
}
