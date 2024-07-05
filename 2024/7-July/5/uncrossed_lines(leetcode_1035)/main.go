package main

func main() {
}

// TOP DOWN
/*
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for row := range dp {
		dp[row] = make([]int, n)
		for col := range n {
			dp[row][col] = -1
		}
	}

	return helper(nums1, nums2, 0, 0, m, n, &dp)
}

func helper(nums1 []int, nums2 []int, i, j, m, n int, dp *[][]int) int {
	if i == m || j == n {
		return 0
	}

	if val := (*dp)[i][j]; val != -1 {
		return val
	}

	var val int
	if nums1[i] == nums2[j] {
		val = helper(nums1, nums2, i+1, j+1, m, n, dp) + 1
	} else {
		val = max(helper(nums1, nums2, i+1, j, m, n, dp), helper(nums1, nums2, i, j+1, m, n, dp))
	}
	(*dp)[i][j] = val
	return val
}
*/
