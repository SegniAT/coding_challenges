package main

import "math"

func main() {

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
