package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
}

func mergeAlternately(word1, word2 string) string {
	result := ""
	if len(word1) < len(word2) {
		for i := 0; i < len(word1); i++ {
			result += word1[i : i+1]
			result += word2[i : i+1]
		}
		result += word2[len(word1):]
	} else if len(word1) == len(word2) {
		for i := 0; i < len(word1); i++ {
			result += word1[i : i+1]
			result += word2[i : i+1]
		}
	} else {
		for i := 0; i < len(word2); i++ {
			result += word1[i : i+1]
			result += word2[i : i+1]
		}
		result += word1[len(word2):]

	}
	return result
}
