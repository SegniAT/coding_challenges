package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{9, 1, 2, 3, 9}
	k := 3

	fmt.Println(largestSumOfAverages(nums, k))
}

// takes a few numbers from the begginning as a subarray then calculates the average and then adds it to the result of the remaining ones
func largestSumOfAverages(nums []int, k int) float64 {
	N := len(nums)
	INF := math.Pow(10, 20)

	cache := make([][]float64, N)
	for i := range cache {
		cache[i] = make([]float64, k+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var helper func(int, int) float64
	helper = func(ind, kLeft int) float64 {
		if ind == N {
			if kLeft == 0 {
				return 0
			}
			return -INF
		}

		if kLeft == 0 {
			return -INF
		}

		if val := cache[ind][kLeft]; val != -1 {
			return val
		}

		best := 0.0
		currentSum := 0
		count := 0

		for i := ind; i < N; i++ {
			currentSum += nums[i]
			count += 1
			best = math.Max(best, float64(currentSum)/float64(count)+helper(i+1, kLeft-1))

		}

		cache[ind][kLeft] = best

		return best
	}

	return helper(0, k)
}
