package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLenNaive(target int, nums []int) int {
	minLen, length := math.Inf(1), len(nums)

	for i := 0; i < length; i++ {
		currSum := 0
		for j := i; j < length; j++ {
			currSum += nums[j]

			if currSum >= target {
				minLen = math.Min(float64(minLen), float64(j-i+1))
				break
			}

		}
	}

	if minLen == math.Inf(1) {
		return 0
	}

	return int(minLen)
}

// Sliding window and Prefix sum
func minSubArrayLen(target int, nums []int) int {
	minLength, length := math.Inf(1), len(nums)

	// construct prefix sum
	prefixSum := make([]int, length)
	prefixSum[0] = nums[0]
	for i := 1; i < length; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	// sliding window algorithm
	left, right := 0, 0

	for right < length && left <= right {
		// get current sum of the window (sum of the numbers between left and right, inclusive)
		currSum := prefixSum[right]
		if left > 0 {
			currSum -= prefixSum[left-1]
		}

		// if current sum is atleast 'target', move the left pointer forward
		// else, move the right pointer forward
		if currSum >= target {
			minLength = math.Min(minLength, float64(right-left+1))
			left++
		} else {
			right++
		}
	}

	if minLength == math.Inf(1) {
		return 0
	}

	return int(minLength)
}
