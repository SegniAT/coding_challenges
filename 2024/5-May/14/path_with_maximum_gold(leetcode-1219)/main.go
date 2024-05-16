package main

import "fmt"

func main() {
	a := [][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}

	fmt.Println(getMaximumGold(a))
}

// up, right, down, left
var DIRS = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func getMaximumGold(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	visited := make([][]bool, rows)
	for row := range rows {
		visited[row] = make([]bool, cols)
	}

	maxGolds := 0
	for row := range rows {
		for col := range cols {
			if grid[row][col] == 0 {
				continue
			}
			golds := dfs(grid, &visited, []int{row, col})
			if maxGolds < golds {
				maxGolds = golds
			}
		}
	}

	return maxGolds
}

func dfs(grid [][]int, visited *[][]bool, coord []int) int {
	maxGold := 0
	for _, dir := range DIRS {
		newCoord := []int{coord[0] + dir[0], coord[1] + dir[1]}
		if inBounds(grid, newCoord) &&
			grid[newCoord[0]][newCoord[1]] != 0 &&
			!(*visited)[newCoord[0]][newCoord[1]] {

			(*visited)[coord[0]][coord[1]] = true
			nxt := dfs(grid, visited, newCoord)
			if nxt > maxGold {
				maxGold = nxt
			}
		}
	}
	(*visited)[coord[0]][coord[1]] = false

	return maxGold + grid[coord[0]][coord[1]]
}

func inBounds(grid [][]int, coord []int) bool {
	return (0 <= coord[0] && coord[0] < len(grid)) && (0 <= coord[1] && coord[1] < len(grid[0]))
}
