package main

import (
	"slices"
)

func main() {
}

func threeSum(nums []int) [][]int {
	N := len(nums)
	res := [][]int{}
	slices.Sort(nums)

	for i := 0; i < N; i++ {
		// to avoid duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// normal 2 sum solution on a sorted array of ints
		j, k := i+1, N-1

		for j < k {
			total := nums[i] + nums[j] + nums[k]

			if total < 0 {
				j++
			} else if total > 0 {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				// move first and check if we've just used the same int above in our solution
				j++

				// to avoid duplicates
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}
		}
	}

	return res
}
