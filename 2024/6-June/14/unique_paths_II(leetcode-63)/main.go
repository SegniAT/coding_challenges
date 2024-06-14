package main

func main() {
}

const OBSTACLE = 1

// TOPDOWN
/*
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == OBSTACLE || obstacleGrid[m-1][n-1] == OBSTACLE {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 1
	helper(obstacleGrid, []int{m - 1, n - 1}, &dp)
	return dp[m-1][n-1]
}

func helper(grid [][]int, cell []int, dp *[][]int) int {
	m, n := len(grid), len(grid[0])
	if !inBounds(cell, m, n) {
		return 0
	}

	row, col := cell[0], cell[1]

	if val := (*dp)[row][col]; val != -1 {
		return val
	}

	if row == 0 && col == 0 {
		return 1
	}

	if grid[row][col] == OBSTACLE {
		return 0
	}

	up := helper(grid, []int{row - 1, col}, dp)
	left := helper(grid, []int{row, col - 1}, dp)

	(*dp)[row][col] = up + left

	return up + left

}
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == OBSTACLE || obstacleGrid[m-1][n-1] == OBSTACLE {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			// top
			if inBounds([]int{row - 1, col}, m, n) && obstacleGrid[row-1][col] != OBSTACLE {
				dp[row][col] += dp[row-1][col]
			}

			// left
			if inBounds([]int{row, col - 1}, m, n) && obstacleGrid[row][col-1] != OBSTACLE {
				dp[row][col] += dp[row][col-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func inBounds(cell []int, m, n int) bool {
	return (0 <= cell[0] && cell[0] < m) && (0 <= cell[1] && cell[1] < n)
}
