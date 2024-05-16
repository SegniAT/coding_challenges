package main

import "fmt"

func main() {
	q := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
	fmt.Println(largestLocal(q))
}

// N, NE, E,
// SE, S, SW,
// W, NW, CENTER
var DIRS = [][]int{
	{-1, 0}, {-1, 1}, {0, 1},
	{1, 1}, {1, 0}, {1, -1},
	{0, -1}, {-1, -1}, {0, 0}}

func largestLocal(grid [][]int) [][]int {
	size := len(grid)

	ans := make([][]int, size-2)
	for i := range size - 2 {
		ans[i] = make([]int, size-2)
	}

	for r := range ans {
		for c := range ans[r] {
			centerX, centerY := r+1, c+1
			fmt.Println(r, c)
			fmt.Println(centerX, centerX)
			max := grid[centerX][centerY]

			for _, dir := range DIRS {
				newX, newY := centerX+dir[0], centerY+dir[1]
				if grid[newX][newY] > max {
					max = grid[newX][newY]
				}
			}

			fmt.Println(max)
			fmt.Println()

			ans[r][c] = max
		}
	}

	return ans
}
