package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	L_MIRROR    = '\\'
	R_MIRROR    = '/'
	V_SPLITTER  = '|'
	H_SPLITTER  = '-'
	EMPTY_SPACE = '.'
)

type Direction rune

const (
	UP    Direction = 'U'
	LEFT  Direction = 'L'
	DOWN  Direction = 'D'
	RIGHT Direction = 'W'
)

var DirectionCoords map[Direction][]int = map[Direction][]int{
	UP:    {-1, 0},
	LEFT:  {0, -1},
	DOWN:  {1, 0},
	RIGHT: {0, 1},
}

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	prettyPrint(input)

	// cell - direction
	visitedCells := make(map[string][]Direction)

	walk(&input, &visitedCells, []int{0, 0}, RIGHT)

	fmt.Println("Energized cells: ", len(visitedCells))
	// Part I
	// First Attempt: 8021 (DONE!!!)

	// Part II

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if _, ok := visitedCells[coordStr([]int{row, col})]; ok {
				input[row][col] = '#'
			} else {
				input[row][col] = '.'
			}
		}
	}

	fmt.Println()

	prettyPrint(input)
}

func walk(matrix *[][]rune, visitedCells *map[string][]Direction, coord []int, dir Direction) {
	if !validCell(matrix, coord) {
		return
	}

	if dirs, ok := (*visitedCells)[coordStr(coord)]; ok {
		if isIn(dir, dirs) {
			return
		} else {
			(*visitedCells)[coordStr(coord)] = append(dirs, dir)
		}
	} else {
		(*visitedCells)[coordStr(coord)] = append(dirs, dir)
	}

	// no need to change direction
	currCellVal := (*matrix)[coord[0]][coord[1]]
	for currCellVal == EMPTY_SPACE ||
		(currCellVal == V_SPLITTER && (dir == UP || dir == DOWN)) ||
		(currCellVal == H_SPLITTER && (dir == LEFT || dir == RIGHT)) {

		switch dir {
		case UP:
			coord[0] += DirectionCoords[UP][0]
			coord[1] += DirectionCoords[UP][1]
		case LEFT:
			coord[0] += DirectionCoords[LEFT][0]
			coord[1] += DirectionCoords[LEFT][1]
		case DOWN:
			coord[0] += DirectionCoords[DOWN][0]
			coord[1] += DirectionCoords[DOWN][1]
		case RIGHT:
			coord[0] += DirectionCoords[RIGHT][0]
			coord[1] += DirectionCoords[RIGHT][1]
		}

		if !validCell(matrix, coord) {
			return
		}

		if dirs, ok := (*visitedCells)[coordStr(coord)]; ok {
			if isIn(dir, dirs) {
				return
			} else {
				(*visitedCells)[coordStr(coord)] = append(dirs, dir)
			}
		} else {
			(*visitedCells)[coordStr(coord)] = append(dirs, dir)
		}

		currCellVal = (*matrix)[coord[0]][coord[1]]
	}

	// change direction
	switch currCellVal {
	case L_MIRROR:
		if dir == UP {
			coord[0] += DirectionCoords[LEFT][0]
			coord[1] += DirectionCoords[LEFT][1]
			walk(matrix, visitedCells, coord, LEFT)
		} else if dir == LEFT {
			coord[0] += DirectionCoords[UP][0]
			coord[1] += DirectionCoords[UP][1]
			walk(matrix, visitedCells, coord, UP)
		} else if dir == DOWN {
			coord[0] += DirectionCoords[RIGHT][0]
			coord[1] += DirectionCoords[RIGHT][1]
			walk(matrix, visitedCells, coord, RIGHT)
		} else if dir == RIGHT {
			coord[0] += DirectionCoords[DOWN][0]
			coord[1] += DirectionCoords[DOWN][1]
			walk(matrix, visitedCells, coord, DOWN)
		}
		return
	case R_MIRROR:
		if dir == UP {
			coord[0] += DirectionCoords[RIGHT][0]
			coord[1] += DirectionCoords[RIGHT][1]
			walk(matrix, visitedCells, coord, RIGHT)
		} else if dir == LEFT {
			coord[0] += DirectionCoords[DOWN][0]
			coord[1] += DirectionCoords[DOWN][1]
			walk(matrix, visitedCells, coord, DOWN)
		} else if dir == DOWN {
			coord[0] += DirectionCoords[LEFT][0]
			coord[1] += DirectionCoords[LEFT][1]
			walk(matrix, visitedCells, coord, LEFT)
		} else if dir == RIGHT {
			coord[0] += DirectionCoords[UP][0]
			coord[1] += DirectionCoords[UP][1]
			walk(matrix, visitedCells, coord, UP)
		}
		return
	case V_SPLITTER:
		if dir == UP {
			coord[0] += DirectionCoords[UP][0]
			coord[1] += DirectionCoords[UP][1]
			walk(matrix, visitedCells, coord, UP)
		} else if dir == LEFT || dir == RIGHT {
			up := []int{coord[0] + DirectionCoords[UP][0], coord[1] + DirectionCoords[UP][1]}
			down := []int{coord[0] + DirectionCoords[DOWN][0], coord[1] + DirectionCoords[DOWN][1]}
			walk(matrix, visitedCells, up, UP)
			walk(matrix, visitedCells, down, DOWN)
		} else if dir == DOWN {
			coord[0] += DirectionCoords[DOWN][0]
			coord[1] += DirectionCoords[DOWN][1]
			walk(matrix, visitedCells, coord, DOWN)
		}
		return
	case H_SPLITTER:
		if dir == UP || dir == DOWN {
			left := []int{coord[0] + DirectionCoords[LEFT][0], coord[1] + DirectionCoords[LEFT][1]}
			right := []int{coord[0] + DirectionCoords[RIGHT][0], coord[1] + DirectionCoords[RIGHT][1]}
			walk(matrix, visitedCells, left, LEFT)
			walk(matrix, visitedCells, right, RIGHT)
		} else if dir == LEFT {
			coord[0] += DirectionCoords[LEFT][0]
			coord[1] += DirectionCoords[LEFT][1]
			walk(matrix, visitedCells, coord, LEFT)
		} else if dir == RIGHT {
			coord[0] += DirectionCoords[RIGHT][0]
			coord[1] += DirectionCoords[RIGHT][1]
			walk(matrix, visitedCells, coord, RIGHT)
		}
		return
	}

}

func isIn(dir Direction, dirs []Direction) bool {
	for _, d := range dirs {
		if dir == d {
			return true
		}
	}

	return false
}

func coordStr(coord []int) string {
	return fmt.Sprintf("%d,%d", coord[0], coord[1])
}

func validCell(matrix *[][]rune, coord []int) bool {
	return (0 <= coord[0] && coord[0] < len(*matrix)) && (0 <= coord[1] && coord[1] < len((*matrix)[0]))
}

func prettyPrint(matrix [][]rune) {
	fmt.Println(strings.Repeat("_", len(matrix[0])*2))
	for _, row := range matrix {
		for _, col := range row {
			fmt.Print(string(col) + " ")
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("_", len(matrix[0])*2))
}

func loadInput(src string) ([][]rune, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	out := [][]rune{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		outRow := []rune{}

		for _, l := range line {
			outRow = append(outRow, rune(l))
		}

		out = append(out, outRow)
	}

	return out, nil
}
