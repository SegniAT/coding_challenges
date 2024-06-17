package main

func main() {
}

func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for row := range rows {
		dp[row] = make([]int, cols)
	}

	maxSquare := 0

	for row := range rows {
		for col := range cols {
			if matrix[row][col] != '1' {
				continue
			}

			if !isInbounds([]int{row - 1, col - 1}, rows, cols) {
				dp[row][col] = 1
			} else {
				dp[row][col] = min(dp[row-1][col-1], min(dp[row-1][col], dp[row][col-1])) + 1
			}

			maxSquare = max(maxSquare, dp[row][col])
		}
	}

	return maxSquare * maxSquare

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isInbounds(cell []int, rows, cols int) bool {
	return (0 <= cell[0] && cell[0] < rows) && (0 <= cell[1] && cell[1] < cols)
}
