package main

import "math"

func main() {
}

// TOP DOWN SOLUTION
/*
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for row := range m {
		dp[row] = make([]int, n)
		for col := range n {
			dp[row][col] = math.MaxInt
		}
	}

	res := helper(dungeon, &dp, []int{0, 0}, m, n)
	return res
}

func helper(dungeon [][]int, dp *[][]int, cell []int, m, n int) int {
	row, col := cell[0], cell[1]

	if !isInbounds(cell, m, n) {
		return math.MaxInt
	}

	if row == m-1 && col == n-1 {
		if dungeon[row][col] <= 0 {
			return 1 - dungeon[row][col]
		} else {
			return 1
		}
	}

	if val := (*dp)[row][col]; val != math.MaxInt {
		return val
	}

	goRight := helper(dungeon, dp, []int{row, col + 1}, m, n)
	goDown := helper(dungeon, dp, []int{row + 1, col}, m, n)

	// if dungeon[row][col] is negative it means it's added and if positive we don't need it so it's subtracted
	// if dungeon[row][col] is too big and the subtraction becomes negative we make it 1 below
	minHealthReq := min(goRight, goDown) - dungeon[row][col]

	if minHealthReq <= 0 {
		minHealthReq = 1
	}

	(*dp)[row][col] = minHealthReq

	return minHealthReq
}

func isInbounds(cell []int, m, n int) bool {
	return (0 <= cell[0] && cell[0] < m) && (0 <= cell[1] && cell[1] < n)
}
*/

// BOTTOM UP SOLUTION
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for row := range m + 1 {
		dp[row] = make([]int, n+1)
		for col := range n + 1 {
			dp[row][col] = math.MaxInt
		}
	}

	// base cases
	dp[m][n-1] = 1
	dp[m-1][n] = 1

	for row := m - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			need := min(dp[row+1][col], dp[row][col+1]) - dungeon[row][col]

			if need <= 0 {
				need = 1
			}

			dp[row][col] = need
		}

	}

	return dp[0][0]
}
