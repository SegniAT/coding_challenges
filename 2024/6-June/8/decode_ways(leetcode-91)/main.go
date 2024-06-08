package main

import "fmt"

func main() {
	fmt.Println(numDecodings("90"))
}

// =======================================================================
// TOP DOWN
// =======================================================================
/*
func numDecodings(s string) int {
	memo := make(map[int]int)
	helper(s, 0, &memo)
	return memo[0]
}

func helper(s string, ind int, memo *map[int]int) int {
	if val, ok := (*memo)[ind]; ok {
		return val
	}

	if ind == len(s) {
		return 1
	}

	ways := 0
	// each character
	if '1' <= s[ind] && s[ind] <= '9' {
		ways += helper(s, ind+1, memo)
	}

	// has at least one more character next to it
	// 10-19...20-26
	if ind+1 < len(s) &&
		((s[ind] == '1' && ('0' <= s[ind+1] && s[ind+1] <= '9')) ||
			(s[ind] == '2' && ('0' <= s[ind+1] && s[ind+1] <= '6'))) {
		ways += helper(s, ind+2, memo)
	}

	(*memo)[ind] = ways
	return ways
}
*/
// =======================================================================
// BOTTOM UP
// =======================================================================

func numDecodings(s string) int {
	ln := len(s)
	dp := make([]int, ln)

	for i := range ln {
		// option one
		if '1' <= s[i] && s[i] <= '9' {
			if i-1 < 0 {
				dp[i] = 1
			} else {
				dp[i] += dp[i-1]
			}
		}

		// option two
		if i-1 >= 0 &&
			((s[i-1] == '1' && '0' <= s[i] && s[i] <= '9') ||
				(s[i-1] == '2' && '0' <= s[i] && s[i] <= '6')) {
			if i-2 < 0 {
				dp[i] = 1
			} else {
				dp[i] += dp[i-2]
			}
		}

	}
	return dp[ln-1]
}
