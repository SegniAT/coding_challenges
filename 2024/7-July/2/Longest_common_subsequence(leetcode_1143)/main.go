package main

import "fmt"

func main() {
	text1 := "a"
	text2 := "a"

	fmt.Println(longestCommonSubsequence(text1, text2))
}

// BOTTOM UP
/*
func longestCommonSubsequence(text1 string, text2 string) int {
	ln1, ln2 := len(text1), len(text2)

	dp := make([][]int, ln1+1)
	for row := range dp {
		dp[row] = make([]int, ln2+1)
	}

	maxLength := 0

	for row := 1; row < len(dp); row++ {
		for col := 1; col < len(dp[0]); col++ {
			if text1[row-1] == text2[col-1] {
				dp[row][col] = dp[row-1][col-1] + 1
			} else {
				dp[row][col] = max(dp[row-1][col], dp[row][col-1])
			}
			maxLength = max(maxLength, dp[row][col])
		}
	}

	return maxLength
}
*/

// TOP DOWN
func longestCommonSubsequence(text1 string, text2 string) int {
	ln1, ln2 := len(text1), len(text2)

	dp := make([][]int, ln1)
	for row := range dp {
		dp[row] = make([]int, ln2)
	}

	for row := range dp {
		for col := range dp[row] {
			dp[row][col] = -1
		}
	}

	maxLength := 0

	helper(&dp, ln1-1, ln2-1, text1, text2)

	for row := range dp {
		for col := range dp[row] {
			maxLength = max(maxLength, dp[row][col])
		}
	}
	return maxLength
}

func helper(dp *[][]int, row, col int, text1, text2 string) int {
	if !isInbounds(*dp, row, col) {
		return 0
	}

	if res := (*dp)[row][col]; res != -1 {
		return res
	}

	var res int

	if text1[row] == text2[col] {
		res = helper(dp, row-1, col-1, text1, text2) + 1
	} else {
		res = max(helper(dp, row-1, col, text1, text2), helper(dp, row, col-1, text1, text2))
	}

	(*dp)[row][col] = res
	return res
}

func isInbounds(grid [][]int, row, col int) bool {
	rows, cols := len(grid), len(grid[0])
	return (0 <= row && row < rows) && (0 <= col && col < cols)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
