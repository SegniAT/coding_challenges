package main

func main() {
}

// up, right, down, left
var DIRS = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

const (
	WATER        = '0'
	LAND         = '1'
	VISITED_LAND = '2'
)

func numIslands(grid [][]byte) int {
	islands := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == LAND {
				markLands([]int{row, col}, &grid)
				islands++
			}
		}
	}
	return islands
}

func markLands(cell []int, grid *[][]byte) {
	row, col := cell[0], cell[1]
	(*grid)[row][col] = VISITED_LAND

	for _, dir := range DIRS {
		newRow, newCol := cell[0]+dir[0], cell[1]+dir[1]
		if isInbounds([]int{newRow, newCol}, *grid) && (*grid)[newRow][newCol] == LAND {
			markLands([]int{newRow, newCol}, grid)
		}
	}
}

func isInbounds(cell []int, grid [][]byte) bool {
	return (0 <= cell[0] && cell[0] < len(grid)) && (0 <= cell[1] && cell[1] < len(grid[0]))
}
