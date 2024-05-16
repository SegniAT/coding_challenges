package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))
}

// main test case to worry about: people=[1, 2, 4, 5] limit=6
func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)

	boatCount := 0
	left, right := 0, len(people)-1

	for left < right {
		if people[left]+people[right] <= limit {
			boatCount++
			left++
		} else {
			boatCount++
		}
		right--
	}

	if left == right {
		boatCount++
	}

	return boatCount
}
