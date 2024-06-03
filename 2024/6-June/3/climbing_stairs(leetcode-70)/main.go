package main

func main() {
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		dp[i+1] += dp[i]
		dp[i+2] += dp[i]
	}

	return dp[n]
}
