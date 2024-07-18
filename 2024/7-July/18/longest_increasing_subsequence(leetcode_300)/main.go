package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))

	nums = []int{0, 1, 0, 3, 2, 3}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)

	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	longest := 0
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
				longest = max(longest, dp[j])
			}
		}
	}

	return longest
}
