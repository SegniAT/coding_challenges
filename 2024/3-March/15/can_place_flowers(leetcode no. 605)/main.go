package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 0}, 1))
}

const EMPTY = 0

func canPlaceFlowers(flowerbed []int, n int) bool {

	for i := range flowerbed {
		if n == 0 {
			return true
		}
		lPossible, rPossible := false, false
		if flowerbed[i] == EMPTY {
			// check left
			if i == 0 || flowerbed[i-1] == EMPTY {
				lPossible = true
			}

			// check right
			if i == len(flowerbed)-1 || flowerbed[i+1] == EMPTY {
				rPossible = true
			}

			if lPossible && rPossible {
				flowerbed[i] = 1
				n--
			}
		}

	}

	return n == 0
}
