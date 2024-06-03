package main

import "fmt"

func main() {
	fmt.Println(countSubarrays([]int{1, 3, 2, 3, 3}, 2))
}

func countSubarrays(nums []int, k int) int64 {
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	var subArrays int64 = 0
	leftPtr := 0
	maxCount := 0

	for _, num := range nums {
		if num == maxNum {
			maxCount++
		}

		for maxCount == k {
			if nums[leftPtr] == maxNum {
				maxCount--
			}
			leftPtr++
		}

		subArrays += int64(leftPtr)
	}

	return subArrays
}
