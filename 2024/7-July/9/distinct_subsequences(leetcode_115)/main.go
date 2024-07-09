package main

import "fmt"

func main() {
	s, t := "abb", "ab"
	s, t = "rabbbit", "rabbit"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)

	dp := make([][]int, sLen)
	for row := range sLen {
		dp[row] = make([]int, tLen)
		for col := range tLen {
			dp[row][col] = -1
		}
	}

	return helper(&dp, s, t, 0, 0, sLen, tLen)
}

func helper(dp *[][]int, s, t string, sInd, tInd, sLen, tLen int) int {
	if sInd >= sLen && tInd >= tLen {
		return 1
	}

	if sInd >= sLen {
		return 0
	}

	if tInd >= tLen {
		return 1
	}

	if res := (*dp)[sInd][tInd]; res != -1 {
		return res
	}

	var res int
	if s[sInd] == t[tInd] {
		res += helper(dp, s, t, sInd+1, tInd, sLen, tLen)
		res += helper(dp, s, t, sInd+1, tInd+1, sLen, tLen)
	} else {
		res += helper(dp, s, t, sInd+1, tInd, sLen, tLen)
	}

	(*dp)[sInd][tInd] = res

	return res
}
