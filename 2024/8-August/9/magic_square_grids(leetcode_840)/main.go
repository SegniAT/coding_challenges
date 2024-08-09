package main

import "fmt"

func main() {
	a := [][]int{
		{5, 5, 5},
		{5, 5, 5},
		{5, 5, 5},
	}

	fmt.Println(isUnique(a, []int{1, 1}))
}

// N,NE,E,SE,S,SW,W,NW
var DIRS = [][2]int{
	{-1, 0}, {-1, 1},
	{0, 1}, {1, 1},
	{1, 0}, {1, -1},
	{0, -1}, {-1, -1},
}

func numMagicSquaresInside(grid [][]int) int {
	N, M := len(grid), len(grid[0])

	magicSquares := 0
	for r := 1; r < N-1; r++ {
		for c := 1; c < M-1; c++ {
			c := []int{r, c}
			if (isValid(grid, c) && isUnique(grid, c)) && hasSameSum(grid, c) {
				magicSquares++
			}
		}

	}

	return magicSquares
}

func isValid(grid [][]int, center []int) bool {
	for _, dir := range DIRS {
		r, c := center[0]+dir[0], center[1]+dir[1]
		val := grid[r][c]
		if !(1 <= val && val <= 9) {
			return false
		}
	}

	return true
}
func isUnique(grid [][]int, center []int) bool {
	cache := make(map[int]bool)
	for _, dir := range DIRS {
		r, c := center[0]+dir[0], center[1]+dir[1]
		val := grid[r][c]
		if _, ok := cache[val]; ok {
			return false
		}
		cache[val] = true
	}
	return true
}

func hasSameSum(grid [][]int, center []int) bool {
	r, N := center[0]-1, center[0]+1
	c, M := center[1]-1, center[1]+1

	// check the rows
	rowsSum := 0
	for j := c; j <= M; j++ {
		rowsSum += grid[r][j]
	}

	for i := r; i <= N; i++ {
		s := 0
		for j := c; j <= M; j++ {
			s += grid[i][j]
		}

		if s != rowsSum {
			return false
		}
	}

	// check the columns
	for i := c; i <= M; i++ {
		s := 0
		for j := r; j <= N; j++ {
			s += grid[j][i]
		}

		if s != rowsSum {
			return false
		}
	}

	// diagonals
	/*
		|  r,c
		|	r+1,c+1
		|		r+2,c+2
	*/
	if grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2] != rowsSum {
		return false
	}

	/*
		|  		r, c+2
		|	r+1,c+1
		| r+2,c
	*/
	if grid[r+2][c]+grid[r+1][c+1]+grid[r][c+2] != rowsSum {
		return false
	}

	return true
}
