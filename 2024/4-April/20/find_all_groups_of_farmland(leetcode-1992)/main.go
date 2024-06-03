package main

import "math"

func main() {
}

const (
	FORESTED_LAND     = 0
	FARM_LAND         = 1
	VISITED_FARM_LAND = 2
)

// up, right, down, left
var DIRS = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func findFarmland(land [][]int) [][]int {
	farmLands := [][]int{}

	for row := range land {
		for col := range land[row] {
			if land[row][col] == FARM_LAND {
				// r1, c1, r2, c2
				group := []int{math.MaxInt, math.MaxInt, math.MinInt, math.MinInt}
				visitLand([]int{row, col}, &group, &land)
				farmLands = append(farmLands, group)
			}
		}
	}

	return farmLands
}

func visitLand(cell []int, group *[]int, land *[][]int) {
	row, col := cell[0], cell[1]
	(*land)[row][col] = VISITED_FARM_LAND

	// group = [r1, c1, r2, c2]
	if row < (*group)[0] {
		(*group)[0] = row
	}

	if col < (*group)[1] {
		(*group)[1] = col
	}

	if row > (*group)[2] {
		(*group)[2] = row
	}

	if col > (*group)[3] {
		(*group)[3] = col
	}

	for _, dir := range DIRS {
		newCell := []int{row + dir[0], col + dir[1]}
		newRow, newCol := newCell[0], newCell[1]
		if isInbounds(newCell, *land) && (*land)[newRow][newCol] == FARM_LAND {
			visitLand(newCell, group, land)
		}
	}
}

func isInbounds(cell []int, grid [][]int) bool {
	row, col := cell[0], cell[1]
	return (0 <= row && row < len(grid)) && (0 <= col && col < len(grid[0]))
}
