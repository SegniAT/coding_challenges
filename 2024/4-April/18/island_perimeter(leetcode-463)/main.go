package main

func main() {
}

// up, right, down, left
var DIRS = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				for _, dir := range DIRS {
					newCell := []int{row + dir[0], col + dir[1]}
					if !isInbounds(newCell, grid) || grid[newCell[0]][newCell[1]] == 0 {
						perimeter++
					}
				}
			}
		}
	}

	return perimeter
}

func isInbounds(cell []int, grid [][]int) bool {
	return (0 <= cell[0] && cell[0] < len(grid)) && (0 <= cell[1] && cell[1] < len(grid[0]))
}
