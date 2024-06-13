package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

// TOP DOWN

/*
func uniquePaths(m int, n int) int {
	dp := [][]int{}
	for range m {
		dp = append(dp, make([]int, n))
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			dp[row][col] = -1
		}
	}

	dp[0][0] = 1

	helper(&dp, []int{m - 1, n - 1}, m, n)

	fmt.Println(dp)

	return dp[m-1][n-1]
}

func helper(dp *[][]int, currCell []int, m, n int) int {
	if !inBounds(currCell, m, n) {
		return 0
	}

	if currCell[0] == 0 && currCell[1] == 0 {
		return (*dp)[0][0]
	}

	if val := (*dp)[currCell[0]][currCell[1]]; val != -1 {
		return val
	}

	up := helper(dp, []int{currCell[0] - 1, currCell[1]}, m, n)
	left := helper(dp, []int{currCell[0], currCell[1] - 1}, m, n)

	(*dp)[currCell[0]][currCell[1]] = up + left

	return up + left
}
*/

// BOTTOM UP
func uniquePaths(m int, n int) int {
	dp := [][]int{}
	for range m {
		dp = append(dp, make([]int, n))
	}

	dp[0][0] = 1

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			// left
			if inBounds([]int{row, col - 1}, m, n) {
				dp[row][col] += dp[row][col-1]
			}

			// up
			if inBounds([]int{row - 1, col}, m, n) {
				dp[row][col] += dp[row-1][col]
			}
		}
	}

	return dp[m-1][n-1]
}

func inBounds(cell []int, m, n int) bool {
	return (0 <= cell[0] && cell[0] < m) && (0 <= cell[1] && cell[1] < n)
}
