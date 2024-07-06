package main

import (
	"fmt"
)

func main() {
	s1 := "sea"
	s2 := "eat"
	fmt.Println(minimumDeleteSum(s1, s2))

	s1 = "bbccacacaab"
	s2 = "aabccb"
	fmt.Println(minimumDeleteSum(s1, s2))
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m)
	for row := range dp {
		dp[row] = make([]int, n)
		for col := range n {
			dp[row][col] = -1
		}
	}

	return helper(&dp, s1, s2, 0, 0, m, n)
}

func helper(dp *[][]int, s1 string, s2 string, i, j, m, n int) int {
	if i >= m {
		return asciiSumFromInd(&s2, j)
	}
	if j >= n {
		return asciiSumFromInd(&s1, i)
	}

	if val := (*dp)[i][j]; val != -1 {
		return val
	}

	var res int
	if s1[i] == s2[j] {
		res = helper(dp, s1, s2, i+1, j+1, m, n)
	} else {
		opt1 := int(s1[i]) + helper(dp, s1, s2, i+1, j, m, n)
		opt2 := int(s2[j]) + helper(dp, s1, s2, i, j+1, m, n)

		res = min(opt1, opt2)
	}

	(*dp)[i][j] = res

	return res
}

func asciiSumFromInd(str *string, i int) int {
	ans := 0
	for ; i < len(*str); i++ {
		ans += int(rune((*str)[i]))
	}
	return ans
}
