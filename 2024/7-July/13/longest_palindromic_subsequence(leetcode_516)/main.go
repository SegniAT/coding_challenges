package main

import "fmt"

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
	s = "cbbd"
	fmt.Println(longestPalindromeSubseq(s))
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for row := range n {
		dp[row] = make([]int, n)
		for col := range n {
			dp[row][col] = -1
		}
	}

	return helper(&dp, s, 0, n-1)
}

func helper(dp *[][]int, s string, i, j int) int {
	if i == j {
		return 1
	}

	if i > j {
		return 0
	}

	if res := (*dp)[i][j]; res != -1 {
		return res
	}

	var res int
	if s[i] == s[j] {
		res = 2 + helper(dp, s, i+1, j-1)
	} else {
		res = max(res, helper(dp, s, i, j-1))
		res = max(res, helper(dp, s, i+1, j))
		res = max(res, helper(dp, s, i+1, j-1))
	}

	(*dp)[i][j] = res
	return res
}
