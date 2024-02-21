package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	/* Use regex
	findWord := regexp.MustCompile(`\b\w+\b`)

	wordMap := findWord.FindAllString(strings.TrimSpace(s), -1)

	if wordMap == nil {
		return s
	}

	for i, j := 0, len(wordMap)-1; i < j; i, j = i+1, j-1 {
		wordMap[i], wordMap[j] = wordMap[j], wordMap[i]
	}

	return strings.Join(wordMap, " ")
	*/

	// Use fields and reverse like normal string

	wordsArr := strings.Fields(s)

	for i, j := 0, len(wordsArr)-1; i < j; i, j = i+1, j-1 {
		wordsArr[i], wordsArr[j] = wordsArr[j], wordsArr[i]
	}
	return strings.TrimSpace(strings.Join(wordsArr, " "))
}
