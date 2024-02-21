package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"),
	)
}

func gcdOfStrings(str1 string, str2 string) string {
	if !strings.Contains(str1, str2) && !strings.Contains(str2, str1) {
		return ""
	}
	result := ""

	if len(str1) > len(str2) {
		for i := 2; i <= len(str2); i++ {
			divideToStr1 := strings.Join(strings.Split(str1, str2[0:i]), "")
			divideToStr2 := strings.Join(strings.Split(str2, str2[0:i]), "")
			if divideToStr1 == "" && divideToStr2 == "" {
				result = str2[0:i]
			}
		}
	} else {
		for i := 2; i <= len(str1); i++ {
			divideToStr2 := strings.Join(strings.Split(str2, str1[0:i]), "")
			divideToStr1 := strings.Join(strings.Split(str1, str1[0:i]), "")
			if divideToStr2 == "" && divideToStr1 == "" {
				result = str1[0:i]
			}

		}
	}

	return result
}
