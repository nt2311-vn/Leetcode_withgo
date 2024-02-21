package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	writeIndex := 0

	for i := 0; i < len(chars); {
		count := 0

		char := chars[i]

		for i < len(chars) && chars[i] == char {
			count++
			i++
		}

		chars[writeIndex] = char
		writeIndex++

		if count > 1 {
			countStr := strconv.Itoa(count)

			for _, c := range countStr {
				chars[writeIndex] = byte(c)
				writeIndex++
			}
		}
	}
	return writeIndex
}
