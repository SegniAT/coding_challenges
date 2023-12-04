package main

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)

}

func rotate(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])

	// transpose
	initRow, initCol := 0, 0
	for initRow < rows && initCol < cols {
		for col := initCol; col < cols; col++ {
			matrix[initRow][col], matrix[col][initRow] = matrix[col][initRow], matrix[initRow][col]
		}

		initRow++
		initCol++
	}

	// flip vertically
	for row := range matrix {
		for col := 0; col < cols/2; col++ {
			oppositeCol := cols - col - 1
			matrix[row][col], matrix[row][oppositeCol] = matrix[row][oppositeCol], matrix[row][col]
		}
	}

}
