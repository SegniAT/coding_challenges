package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 1, 2, 3, 1, 2}
	k := 2
	/*
			fmt.Println(maxSubarrayLength(nums, k))


			nums = []int{5, 5, 5, 5, 5, 5, 5}
			k = 4

			fmt.Println(maxSubarrayLength(nums, k))

		nums = []int{1, 1, 1, 1, 2}
		k = 2

		fmt.Println(maxSubarrayLength(nums, k))
	*/

	nums = []int{1, 2, 1, 2, 1, 2, 1, 2}
	k = 1

	fmt.Println(maxSubarrayLength(nums, k))
}

func maxSubarrayLength(nums []int, k int) int {
	leftPtr := 0
	countMap := make(map[int]int)

	maxLen := 0

	for rightPtr, num := range nums {
		countMap[num]++

		for countMap[num] > k {
			countMap[nums[leftPtr]]--
			leftPtr++
		}

		if newLen := rightPtr - leftPtr + 1; newLen > maxLen {
			maxLen = newLen
		}

	}

	return maxLen
}
