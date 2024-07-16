package main

func main() {
}

func countSubstrings(s string) int {
	n := len(s)

	dp := make([][]int, n)
	for row := range n {
		dp[row] = make([]int, n)
		for col := range n {
			dp[row][col] = -1
		}
	}

	var helper func(int, int) int

	helper = func(i, j int) int {
		if i >= j {
			return 1
		}

		if res := dp[i][j]; res != -1 {
			return res
		}

		res := 0
		if s[i] == s[j] {
			res = helper(i+1, j-1)
		}

		dp[i][j] = res

		return res
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			res += helper(i, j)
		}
	}

	return res

}
