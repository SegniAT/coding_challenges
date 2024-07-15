package main

import ()

func main() {
}

// why does subtracting work?? - because the consecutive helper calls in the stack are alice and bob
// why take the maximum of the two options ??

func stoneGameVII(stones []int) int {
	n := len(stones)
	prefixSum := make([]int, n+1)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + stones[i]
		dp[i] = make([]int, n)
	}

	var helper func(int, int) int
	helper = func(i, j int) int {
		if i >= j {
			return 0
		}

		var res int
		if res = dp[i][j]; res > 0 {
			return res
		}

		// pick out leftmost stone
		opt1 := prefixSum[j+1] - prefixSum[i+1] - helper(i+1, j)

		// pick out rightmost stone
		opt2 := prefixSum[j] - prefixSum[i] - helper(i, j-1)

		if opt1 > opt2 {
			dp[i][j] = opt1
			return opt1
		}

		dp[i][j] = opt2
		return opt2
	}

	return helper(0, n-1)
}
