package main

func main() {
}

func tribonacci(n int) int {
	dp := make([]int, n+10)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 0; i+3 < n+10; i++ {
		dp[i+3] = dp[i] + dp[i+1] + dp[i+2]
	}

	return dp[n]
}
