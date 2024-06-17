package main

func main() {
}

func countSquares(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for row := range rows {
		dp[row] = make([]int, cols)
	}

	count := 0

	for row := range rows {
		for col := range cols {
			if matrix[row][col] != 1 {
				continue
			}

			if !isInbounds([]int{row - 1, col - 1}, rows, cols) {
				dp[row][col] = 1
			} else {
				dp[row][col] = min(dp[row-1][col-1], min(dp[row-1][col], dp[row][col-1])) + 1
			}

			count += dp[row][col]
		}
	}

	return count
}

func isInbounds(cell []int, rows, cols int) bool {
	return (0 <= cell[0] && cell[0] < rows) && (0 <= cell[1] && cell[1] < cols)
}
