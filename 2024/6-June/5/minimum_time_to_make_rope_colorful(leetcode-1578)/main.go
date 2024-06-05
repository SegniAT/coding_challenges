package main

import (
	"fmt"
)

func main() {
	colors := "bbbaaa"
	neededTime := []int{4, 9, 3, 8, 8, 9}

	fmt.Println(minCost(colors, neededTime))
}

// delete every color in a contiguous subarray of same colors except the maximum
func minCost(colors string, neededTime []int) int {
	var maxSameColorCost, runningSum, minCost int
	prev := ' '

	for ind, color := range colors {
		if color != prev {
			minCost += runningSum - maxSameColorCost
			maxSameColorCost, runningSum = 0, 0
		}

		if neededTime[ind] > maxSameColorCost {
			maxSameColorCost = neededTime[ind]
		}
		prev = color
		runningSum += neededTime[ind]
	}

	// possible continious subarray found at the end of the array
	minCost += runningSum - maxSameColorCost
	return minCost
}
