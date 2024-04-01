package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}

// change 0's into -1, the problem becomes find the maximum sub-array with the prefix sum of 0

func findMaxLength(nums []int) int {
	res, sum := 0, 0
	prev := make(map[int]int)

	for i, num := range nums {
		if num == 0 {
			sum--
		} else {
			sum++
		}

		if sum == 0 {
			res = i + 1
		} else if prevInd, ok := prev[sum]; !ok {
			prev[sum] = i
		} else {
			res = int(math.Max(float64(res), float64(i-prevInd)))
		}

	}

	return res
}
