package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Direction struct {
	from rune
	diff []int
}

var dirs = map[rune]Direction{
	'N': {
		from: 'S',
		diff: []int{-1, 0},
	},
	'E': {
		from: 'W',
		diff: []int{0, 1},
	},
	'S': {
		from: 'N',
		diff: []int{1, 0},
	},
	'W': {
		from: 'E',
		diff: []int{0, -1},
	},
}

var tilesDir = map[rune][]rune{
	'|': {'N', 'S'},
	'-': {'E', 'W'},
	'L': {'N', 'E'},
	'J': {'N', 'W'},
	'7': {'W', 'S'},
	'F': {'E', 'S'},
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	rows, cols := len(lines), len(lines[0])

	input := make([][]rune, rows)
	start := []int{0, 0}

	for ind, line := range lines {
		input[ind] = make([]rune, cols)
		for i := 0; i < cols; i++ {
			char := rune(line[i])
			input[ind][i] = char
			if char == 'S' {
				start = []int{ind, i}
			}

		}
	}

	// fmt.Println(tiles)
	// fmt.Println(input)
	// fmt.Println(start)

	startNeighbours := getValidNeighbours(input, start)
	distances := make(map[string]int)

	for _, startNeighbour := range startNeighbours {
		currCoord := startNeighbour
		prevDir := '.'
		// find from where we got here
		for _, dir := range dirs {
			nextCoord := addCoords(start, dir.diff)
			if sameCoords(currCoord, nextCoord) {
				prevDir = dir.from
				break
			}
		}

		distances[coordToString(currCoord)] = 1

		for !sameCoords(start, currCoord) {
			// fmt.Println(currCoord, string(input[currCoord[0]][currCoord[1]]))
			currTile := input[currCoord[0]][currCoord[1]]
			currTileDirs := tilesDir[currTile]

			var nextDir Direction

			if currTileDirs[0] == prevDir {
				nextDir = dirs[currTileDirs[1]]
			} else {
				nextDir = dirs[currTileDirs[0]]
			}

			nextCoord := addCoords(currCoord, nextDir.diff)
			// check if the next coordinate is valid
			if !isInRange(nextCoord, rows, cols) {
				break
			}

			nextTile := input[nextCoord[0]][nextCoord[1]]

			// check if the tile exists in map
			if _, ok := tilesDir[nextTile]; !ok {
				break
			}

			nextTileDirs := tilesDir[nextTile]

			// check if the next tile is connected to the current
			if !(nextDir.from != nextTileDirs[0] || nextDir.from != nextTileDirs[1]) {
				break
			}

			nextCoordStr := coordToString(nextCoord)
			currDistance := distances[coordToString(currCoord)]
			if _, ok := distances[nextCoordStr]; ok {
				distances[nextCoordStr] = int(math.Min(float64(currDistance)+1, float64(distances[nextCoordStr])))
			} else {
				distances[nextCoordStr] = currDistance + 1
			}

			prevDir = nextDir.from
			currCoord = nextCoord
		}

		// fmt.Println()
		// fmt.Println()
		// fmt.Println()
		// fmt.Println()
		// fmt.Println(distances)
	}

	maxDistance := -1

	for _, val := range distances {
		if val > maxDistance {
			maxDistance = val
		}
	}

	fmt.Println("Max distance: ", maxDistance)

}

func addCoords(coord1, coord2 []int) []int {
	return []int{coord1[0] + coord2[0], coord1[1] + coord2[1]}
}

func sameCoords(coord1, coord2 []int) bool {
	return coord1[0] == coord2[0] && coord1[1] == coord2[1]
}

func getValidNeighbours(input [][]rune, start []int) [][]int {
	out := [][]int{}

	for _, dir := range dirs {
		nextCoord := []int{start[0] + dir.diff[0], start[1] + dir.diff[1]}
		if !isInRange(nextCoord, len(input), len(input[0])) {
			continue
		}

		nextTile := input[nextCoord[0]][nextCoord[1]]
		if _, ok := tilesDir[nextTile]; !ok {
			continue
		}

		if !(dir.from == tilesDir[nextTile][0] || dir.from == tilesDir[nextTile][1]) {
			continue
		}

		out = append(out, nextCoord)
	}

	return out
}

func coordToString(coord []int) string {
	return fmt.Sprintf("%d,%d", coord[0], coord[1])
}

func isInRange(coord []int, rows, cols int) bool {
	return (0 <= coord[0] && coord[0] < rows) && (0 <= coord[1] && coord[1] < cols)
}

func readInput(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := []string{}

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, nil
}
