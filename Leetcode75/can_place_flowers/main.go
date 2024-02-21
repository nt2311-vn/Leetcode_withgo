package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	/*
		If else normal methof
		if len(flowerbed) == 1 {
			return flowerbed[0] == 0
		}
		if n == 0 {
			return true
		}
		count := 0
		for i := 0; i < len(flowerbed); i++ {
			if i == 0 {
				if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
					count++
					flowerbed[i] = 1
				}
				if count >= n {
					return true
				}
				continue
			}

			if i == len(flowerbed)-1 {
				if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
					count++
					flowerbed[i] = 1
				}
				if count >= n {
					return true
				}
				continue
			}

			if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				count++
				flowerbed[i] = 1

				if count >= n {
					return true
				}
			}
		}
		return count >= n
	*/
	// Try implement recursive function

	if n == 0 {
		return true
	}

	if len(flowerbed) == 0 {
		return false
	}

	if flowerbed[0] == 0 && (len(flowerbed) == 1 || flowerbed[1] == 0) {
		// Place a flower at the current position
		updatedFlowerbed := append([]int{1}, flowerbed[2:]...)
		return canPlaceFlowers(updatedFlowerbed, n-1)
	}

	// Move to the next position without placing a flower
	return canPlaceFlowers(flowerbed[1:], n)
}
