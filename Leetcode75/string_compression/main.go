package main

import (
	"fmt"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	mapCompress := map[byte]int{}

	for _, v := range chars {
		_, exist := mapCompress[v]

		if exist {
			mapCompress[v]++
		} else {
			mapCompress[v] = 1
		}
	}

	chars = make([]byte, 0)

	for key, value := range mapCompress {
		chars = append(chars, key)
		chars = append(chars, byte(value))
	}

	fmt.Println(chars)

	return len(chars)
}
