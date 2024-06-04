package main

import "math"

func main() {
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		option1 := dp[i-1] + cost[i-1]
		option2 := dp[i-2] + cost[i-2]

		dp[i] = int(math.Min(float64(option1), float64(option2)))
	}

	return dp[n]
}
