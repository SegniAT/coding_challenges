package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction rune

const (
	UP    Direction = 'U'
	RIGHT Direction = 'R'
	DOWN  Direction = 'D'
	LEFT  Direction = 'L'

	TRENCH     = '#'
	GL_TERRAIN = '.'
)

type DigInfo struct {
	Direction Direction
	Length    int
	Hex       string
}

var DirCoords map[Direction][]int = map[Direction][]int{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

func main() {
	digPlan, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	rows, cols := 0, 0
	r, c := 0, 0
	minRow, maxRow := math.MaxInt, math.MinInt
	minCol, maxCol := math.MaxInt, math.MinInt

	for _, plan := range digPlan {
		switch plan.Direction {
		case DOWN:
			r += plan.Length
			if r > maxRow {
				maxRow = r
			}
		case UP:
			r -= plan.Length
			if r < minRow {
				minRow = r
			}
		case RIGHT:
			c += plan.Length
			if c > maxCol {
				maxCol = c
			}
		case LEFT:
			c -= plan.Length
			if c < minCol {
				minCol = c
			}
		}
	}

	rows, cols = maxRow-minRow+1, maxCol-minCol+1
	fmt.Println("minRow maxRow: ", minRow, maxRow)
	fmt.Println("minCol maxCol: ", minCol, maxCol)
	fmt.Println("rows cols: ", rows, cols)

	topView := make([][]rune, rows)

	for row := 0; row < rows; row++ {
		topView[row] = make([]rune, cols)
		for col := 0; col < cols; col++ {
			topView[row][col] = GL_TERRAIN
		}
	}

	walk(topView, []int{int(math.Abs(float64(minRow))), int(math.Abs(float64(minCol)))}, digPlan)

	// prettyPrint(topView)

	for row := 0; row < rows; row++ {
		paint := false
		for col := 0; col < cols; {
			char := topView[row][col]

			if char == TRENCH {
				paint = !paint
				for col < cols && topView[row][col] == TRENCH {
					col++
				}
				continue
			} else if char == GL_TERRAIN && paint {
				topView[row][col] = TRENCH
			}

			col++
		}
	}

	// prettyPrint(topView)

	cubicMeters := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if topView[row][col] == TRENCH {
				cubicMeters++
			}
		}
	}

	fmt.Println("Cubic meters: ", cubicMeters)
	// PART I
	// Attemp 1: 48414 (TOO HIGH)
}

func walk(matrix [][]rune, coord []int, digPlan []DigInfo) {
	matrix[coord[0]][coord[1]] = TRENCH

	for _, dplan := range digPlan {
		for i := 0; i < dplan.Length; i++ {
			coord = addCoords(coord, DirCoords[dplan.Direction])
			matrix[coord[0]][coord[1]] = TRENCH
		}
	}
}

func isValidCoord(matrix [][]any, coord []int) bool {
	return (0 <= coord[0] && coord[0] < len(matrix)) && (0 <= coord[1] && coord[1] < len(matrix[0]))
}

func prettyPrint(matrix [][]rune) {
	fmt.Println(strings.Repeat("_", len(matrix[0])*3))
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf(" %c ", col)
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("_", len(matrix[0])*3))
}

func addCoords(coord1, coord2 []int) []int {
	return []int{coord1[0] + coord2[0], coord1[1] + coord2[1]}
}

func loadInput(src string) ([]DigInfo, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := []DigInfo{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		dir := Direction(rune(line[0][0]))
		length, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		color := line[2][1 : len(line[2])-1]

		out = append(out, DigInfo{
			Direction: dir,
			Length:    length,
			Hex:       color,
		})
	}

	return out, nil
}
