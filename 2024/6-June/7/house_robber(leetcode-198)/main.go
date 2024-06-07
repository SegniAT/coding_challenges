package main

func main() {
}

// =================================================================================================
// TOP DOWN SOLUTION
// =================================================================================================

/*
func rob(nums []int) int {
	ln := len(nums)
	dp := []int{}
	for range ln {
		dp = append(dp, -1)
	}

	a, b := dfs(ln-1, nums, &dp), dfs(ln-2, nums, &dp)
	return max(a, b)
}

func dfs(index int, nums []int, dp *[]int) int {
	if index < 0 {
		return 0
	}

	if cached := (*dp)[index]; cached != -1 {
		return cached
	}

	opt1 := dfs(index-2, nums, dp)
	opt2 := dfs(index-3, nums, dp)

	res := max(opt1, opt2) + nums[index]
	(*dp)[index] = res
	return res
}
*/

// =================================================================================================
// BOTTOM UP SOLUTION
// =================================================================================================

func rob(nums []int) int {
	ln := len(nums)
	if ln == 1 {
		return nums[0]
	}

	dp := []int{}
	for i := range ln {
		dp = append(dp, nums[i])
	}

	for i := range ln {
		opt1 := 0
		if i-2 >= 0 {
			opt1 = dp[i-2]
		}

		opt2 := 0
		if i-3 >= 0 {
			opt2 = dp[i-3]
		}

		dp[i] += max(opt1, opt2)
	}

	return max(dp[ln-1], dp[ln-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
