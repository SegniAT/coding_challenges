package main

import "fmt"

func main() {
	word1 := "intention"
	word2 := "execution"

	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m)
	for row := range dp {
		dp[row] = make([]int, n)
		for col := range n {
			dp[row][col] = -1
		}
	}

	return helper(&dp, word1, word2, 0, 0, m, n)
}

func helper(dp *[][]int, word1, word2 string, ind1, ind2, len1, len2 int) int {
	if ind1 >= len1 {
		return countForward(ind2, len2)
	}

	if ind2 >= len2 {
		return countForward(ind1, len1)
	}

	if word1[ind1] == word2[ind2] {
		return helper(dp, word1, word2, ind1+1, ind2+1, len1, len2)
	}

	if val := (*dp)[ind1][ind2]; val != -1 {
		return val
	}

	// insert
	opt1 := 1 + helper(dp, word1, word2, ind1, ind2+1, len1, len2)

	// delete
	opt2 := 1 + helper(dp, word1, word2, ind1+1, ind2, len1, len2)

	//replace
	opt3 := 1 + helper(dp, word1, word2, ind1+1, ind2+1, len1, len2)

	res := min(opt1, opt2, opt3)
	(*dp)[ind1][ind2] = res

	return res
}

func countForward(i int, ln int) int {
	count := 0
	for ; i < ln; i++ {
		count++
	}
	return count
}
