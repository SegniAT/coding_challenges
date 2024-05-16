package main

import (
	"fmt"
	"math"
)

func main() {
	a := [][]int{
		{0, 0, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
	}

	fmt.Println(matrixScore(a))
}

var INDEX_OUT_OF_BOUNDS_ERR = fmt.Errorf("Index out of bounds")
var WRONG_CELL_VALUE_ERR = fmt.Errorf("Wrong cell value")

func matrixScore(grid [][]int) int {
	// handle the rows
	for row := range grid {
		if grid[row][0] == 0 {
			flipRow(&grid, row)
		}
	}

	rows, cols := len(grid), len(grid[0])
	for col := range cols {
		zeroCount := 0
		for row := range rows {
			if grid[row][col] == 0 {
				zeroCount++
			}
		}

		if zeroCount > rows/2 {
			flipCol(&grid, col)
		}
	}

	ans := 0

	for _, row := range grid {
		ans += bitArrayToNum(row)
	}

	return ans
}

func flipRow(grid *[][]int, row int) error {
	rows, cols := len(*grid), len((*grid)[0])
	if row >= rows {
		return INDEX_OUT_OF_BOUNDS_ERR
	}

	for col := range cols {
		oldVal := (*grid)[row][col]

		switch oldVal {
		case 0:
			(*grid)[row][col] = 1
		case 1:
			(*grid)[row][col] = 0
		default:
			return WRONG_CELL_VALUE_ERR
		}
	}

	return nil
}

func flipCol(grid *[][]int, col int) error {
	rows, cols := len(*grid), len((*grid)[0])
	if col >= cols {
		return INDEX_OUT_OF_BOUNDS_ERR
	}

	for row := range rows {
		oldVal := (*grid)[row][col]

		switch oldVal {
		case 0:
			(*grid)[row][col] = 1
		case 1:
			(*grid)[row][col] = 0
		default:
			return WRONG_CELL_VALUE_ERR
		}
	}

	return nil
}

func bitArrayToNum(arr []int) int {
	num := 0
	maxInd := len(arr) - 1

	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			num += int(math.Pow(2, float64(maxInd-i)))
		}
	}

	return num
}
