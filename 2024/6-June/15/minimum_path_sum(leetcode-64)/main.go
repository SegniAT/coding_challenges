package main

import "math"

func main() {
}

// TOP DOWN

/*
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for row := 0; row < m; row++ {
		dp[row] = make([]int, n)
		for col := 0; col < n; col++ {
			dp[row][col] = math.MaxInt
		}
	}
	dp[0][0] = grid[0][0]
	helper(grid, &dp, []int{m - 1, n - 1}, m, n)
	return dp[m-1][n-1]

}

func helper(grid [][]int, dp *[][]int, cell []int, m, n int) int {
	row, col := cell[0], cell[1]
	if !inBounds(cell, m, n) {
		return math.MaxInt
	}

	if row == 0 && col == 0 {
		return grid[0][0]
	}

	if val := (*dp)[row][col]; val != math.MaxInt {
		return val
	}

	left := helper(grid, dp, []int{row, col - 1}, m, n)
	up := helper(grid, dp, []int{row - 1, col}, m, n)

	(*dp)[row][col] = min(left, up) + grid[row][col]

	return min(left, up) + grid[row][col]
}
*/

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for row := 0; row < m; row++ {
		dp[row] = make([]int, n)
		for col := 0; col < n; col++ {
			dp[row][col] = math.MaxInt
		}
	}
	dp[0][0] = grid[0][0]

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			left, up := math.MaxInt, math.MaxInt

			// left
			if inBounds([]int{row, col - 1}, m, n) {
				left = dp[row][col-1]
			}

			// up
			if inBounds([]int{row - 1, col}, m, n) {
				up = dp[row-1][col]
			}

			if left == math.MaxInt && up == math.MaxInt {
				left, up = 0, 0
			}

			dp[row][col] = min(left, up) + grid[row][col]
		}

	}

	return dp[m-1][n-1]
}

func inBounds(cell []int, m, n int) bool {
	return (0 <= cell[0] && cell[0] < m) && (0 <= cell[1] && cell[1] < n)
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
