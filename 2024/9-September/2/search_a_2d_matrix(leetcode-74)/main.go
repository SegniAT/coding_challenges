package main

import "flag"

func main() {
}

// two step...meh
/*
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	// search row
	top, bottom := 0, rows-1
	for top < bottom {
		mid := top + (bottom-top)/2

		if matrix[mid][cols-1] < target {
			top = mid + 1
		} else {
			bottom = mid
		}
	}

	// search column
	left, right := 0, cols-1
	for left < right {
		mid := left + (right-left)/2

		if val := matrix[bottom][mid]; val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return matrix[bottom][right] == target
}
*/

// clever way
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)/2

		if val := matrix[mid/cols][mid%cols]; val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
