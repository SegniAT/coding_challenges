package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(maximumHappinessSum([]int{12, 1, 42}, 3))
}

// TLE
func maximumHappinessSum1(happiness []int, k int) int64 {
	var sum int64

	slices.Sort(happiness)
	slices.Reverse(happiness)

	for k > 0 && len(happiness) > 0 {
		if happiness[0] <= 0 {
			break
		}
		sum += int64(happiness[0])
		happiness = happiness[1:]

		for i := range happiness {
			happiness[i]--
		}

		k--
	}

	return sum
}

func maximumHappinessSum(happiness []int, k int) int64 {
	var sum int64

	slices.Sort(happiness)
	slices.Reverse(happiness)

	turn := 0
	biggestInd := 0

	ln := len(happiness)

	for k > 0 && biggestInd < ln {
		newNum := happiness[biggestInd] - turn
		if newNum <= 0 {
			break
		}

		sum += int64(newNum)

		k--
		turn++
		biggestInd++
	}

	return sum
}
