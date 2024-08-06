package main

import "fmt"

func main() {
}

// must include every element...any element can be + or -
func findTargetSumWays(nums []int, target int) int {
	N := len(nums)
	cache := make(map[string]int)
	var helper func(int, int) int
	helper = func(targetLeft int, ind int) int {
		if ind < 0 {
			if targetLeft == 0 {
				return 1
			}
			return 0
		}

		key := fmt.Sprintf("%d,%d", targetLeft, ind)
		if val, ok := cache[key]; ok {
			return val
		}

		opt1 := helper(targetLeft+nums[ind], ind-1)
		opt2 := helper(targetLeft-nums[ind], ind-1)

		cache[key] = opt1 + opt2

		return opt1 + opt2
	}

	return helper(target, N-1)
}
