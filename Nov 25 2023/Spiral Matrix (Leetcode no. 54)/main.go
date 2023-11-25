package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	res := []int{}

	initRow, lastRow := 0, len(matrix)-1
	initCol, lastCol := 0, len(matrix[0])-1
	elements := len(matrix) * len(matrix[0])

	for len(res) < elements {
		// first row : left to right
		for i := initCol; i <= lastCol; i++ {
			res = append(res, matrix[initRow][i])
		}
		initRow++

		if len(res) == elements {
			break
		}

		// last column : top to bottom
		for i := initRow; i <= lastRow; i++ {
			res = append(res, matrix[i][lastCol])
		}
		lastCol--

		if len(res) == elements {
			break
		}

		// last row : right to left
		for i := lastCol; i >= initCol; i-- {
			res = append(res, matrix[lastRow][i])
		}
		lastRow--

		if len(res) == elements {
			break
		}

		// first column : bottom to top
		for i := lastRow; i >= initRow; i-- {
			res = append(res, matrix[i][initCol])
		}
		initCol++
	}

	return res
}
