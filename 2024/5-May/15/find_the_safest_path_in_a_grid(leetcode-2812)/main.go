package main

import (
	"fmt"
	"math"
)

func main() {
	inp := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}

	fmt.Println(maximumSafenessFactor(inp))

	inp = [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{0, 0, 0},
	}

	fmt.Println(maximumSafenessFactor(inp))

	inp = [][]int{
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
	}

	fmt.Println(maximumSafenessFactor(inp))
}

const (
	EMPTY = 0
	THIEF = 1
)

// up, right, down, left
var DIRS = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func maximumSafenessFactor(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	distance := make([][]int, rows)
	for row := range rows {
		distance[row] = make([]int, cols)
	}

	// init to negative
	for row := range rows {
		for col := range cols {
			distance[row][col] = -1
		}
	}

	for row := range rows {
		for col := range cols {
			if grid[row][col] == 1 {
				bfs(grid, &distance, []int{row, col})
			}
		}
	}

	fmt.Println(distance)

	return -1
}

// distance of thiefs to every cell
func bfs(grid [][]int, distance *[][]int, cell []int) {
	rows, cols := len(grid), len(grid[0])

	visited := make([][]bool, rows)
	for row := range rows {
		visited[row] = make([]bool, cols)
	}

	visited[cell[0]][cell[1]] = true

	queue := [][]int{cell}
	dist := 0

	size := len(queue)
	for size > 0 {
		for i := 0; i < size; i++ {
			// pop
			currCell := queue[0]
			queue = queue[1:]

			visited[currCell[0]][currCell[1]] = true

			if old := (*distance)[currCell[0]][currCell[1]]; old == -1 {
				(*distance)[currCell[0]][currCell[1]] = dist
			} else {
				(*distance)[currCell[0]][currCell[1]] = int(math.Min(float64(old), float64(dist)))
			}

			for _, dir := range DIRS {
				newCell := []int{dir[0] + currCell[0], dir[1] + currCell[1]}
				if inBounds(grid, newCell) &&
					!visited[newCell[0]][newCell[1]] &&
					grid[newCell[0]][newCell[1]] == EMPTY {
					queue = append(queue, newCell)
				}
			}
		}
		size = len(queue)
		dist++
	}
}

func inBounds(grid [][]int, cell []int) bool {
	row, col := cell[0], cell[1]
	return (0 <= row && row < len(grid)) && (0 <= col && col < len(grid[0]))
}
