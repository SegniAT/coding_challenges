package main

func main() {

}

func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	rowIsZero := make([]bool, rows)
	colIsZero := make([]bool, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				rowIsZero[row] = true
				colIsZero[col] = true
			}
		}
	}

	for row := rows - 1; row > -1; row-- {
		for col := cols - 1; col > -1; col-- {
			if rowIsZero[row] || colIsZero[col] {
				matrix[row][col] = 0
			}
		}
	}
}
