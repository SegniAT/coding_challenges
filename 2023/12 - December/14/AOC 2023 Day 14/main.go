package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EMPTY        = '.'
	ROUNDED_ROCK = 'O'
	CUBE_ROCK    = '#'
)

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	// fmt.Println("Input: ")
	// prettyPrint(input)
	// fmt.Println()

	ohLord := make(map[string]bool)

	matrixStr := matrixToString(input)
	ohLord[matrixStr] = true
	// prev := copyMatrix(input)
	firstThousand := []int{}
	for i := 0; i < 1_000; i++ {
		rollNorth(&input)
		rollWest(&input)
		rollSouth(&input)
		rollEast(&input)

		// fmt.Println("Cycle: ", i, " ", getLoad(input))
		// matrixStr = matrixToString(input)
		// if _, ok := ohLord[matrixStr]; ok {
		// 	break
		// } else {
		// 	prev = copyMatrix(input)
		// 	ohLord[matrixStr] = true
		// }
		// prettyPrint(input)
		load := getLoad(input)
		firstThousand = append(firstThousand, load)
		// fmt.Printf("%d: %d\n", i, load)
		// fmt.Println()
	}

	fmt.Println("TOTAL LOAD: ", getLoad(input))
	cycleHeadInd := -1
	for i := 0; i < len(firstThousand); i++ {
		if firstThousand[i] == 94491 {
			cycleHeadInd = i
			break
		}
	}

	cycle := []int{94491, 94502, 94511, 94535, 94560, 94591, 94618, 94620, 94615, 94585, 94565, 94537, 94510, 94493, 94487, 94494, 94517, 94536, 94565, 94587, 94610, 94626, 94616, 94590, 94561, 94529, 94516, 94494, 94492, 94490, 94509, 94542, 94566, 94592, 94606, 94618, 94622, 94591, 94566, 94525, 94508, 94500, 94493, 94495, 94505, 94534, 94572, 94593, 94611, 94614, 94614, 94597, 94567, 94530, 94504, 94492, 94499, 94496, 94510, 94530, 94564, 94599, 94612, 94619, 94610, 94589, 94573, 94531, 94509, 94488}
	cycleLen := len(cycle)

	fmt.Println("1, 000: ", cycle[(1_000-cycleHeadInd)%cycleLen])
	fmt.Println("1, 000, 000, 000: ", cycle[(1_000_000_000-cycleHeadInd)%cycleLen])

	// OFF BY ONE ERROR ABOVE, DO A -1 on the index

	// stupid way
	j := 0
	for i := cycleHeadInd; i < 1_000_000_000; i++ {
		j++
		if j >= cycleLen {
			j = 0
		}
	}

	fmt.Println(j, cycle[j])

	// Part I
	// Attempt 1: 109466 (DONE!!)

	// Part II
	// Attempt 1: 94565 (too low)
	// Attempt 2: 94585 (RIGHT ANSWER!!!)
}

func getLoad(matrix [][]rune) int {
	rows, cols := len(matrix), len(matrix[0])
	load := 0

	for row := 0; row < rows; row++ {
		rowLoad, roundRockCount := rows-row, 0
		for col := 0; col < cols; col++ {
			currChar := matrix[row][col]
			if currChar == ROUNDED_ROCK {
				roundRockCount++
			}
		}

		load += (rowLoad * roundRockCount)
	}

	return load
}

func copyMatrix(input [][]rune) [][]rune {
	out := [][]rune{}

	for _, row := range input {
		newRow := []rune{}
		for _, col := range row {
			newRow = append(newRow, col)
		}
		out = append(out, newRow)
	}

	return out
}

func matrixToString(matrix [][]rune) string {
	out := ""
	for _, row := range matrix {
		rowStr := ""
		for _, col := range row {
			rowStr = fmt.Sprintf("%s%c", rowStr, col)
		}
		out = fmt.Sprintf("%s(%s)", out, rowStr)
	}
	return out
}

func rollNorth(matrix *[][]rune) {
	rows, cols := len(*matrix), len((*matrix)[0])
	for col := 0; col < cols; col++ {
		availableSlot := -1
		for row := 0; row < rows; row++ {
			currChar := (*matrix)[row][col]

			switch currChar {
			case EMPTY:
				if availableSlot == -1 {
					availableSlot = row
				}
			case ROUNDED_ROCK:
				if availableSlot != -1 {
					(*matrix)[row][col] = EMPTY
					(*matrix)[availableSlot][col] = ROUNDED_ROCK
					availableSlot++
					if (*matrix)[availableSlot][col] == CUBE_ROCK {
						availableSlot = -1
					}
				}

			case CUBE_ROCK:
				availableSlot = -1
			}
		}
	}
}

func rollWest(matrix *[][]rune) {
	rows, cols := len(*matrix), len((*matrix)[0])

	for row := 0; row < rows; row++ {
		availableSlot := -1
		for col := 0; col < cols; col++ {
			currChar := (*matrix)[row][col]

			switch currChar {
			case EMPTY:
				if availableSlot == -1 {
					availableSlot = col
				}
			case ROUNDED_ROCK:
				if availableSlot != -1 {
					(*matrix)[row][col] = EMPTY
					(*matrix)[row][availableSlot] = ROUNDED_ROCK
					availableSlot++
					if (*matrix)[row][availableSlot] == CUBE_ROCK {
						availableSlot = -1
					}
				}
			case CUBE_ROCK:
				availableSlot = -1
			}
		}
	}
}

func rollSouth(matrix *[][]rune) {
	rows, cols := len(*matrix), len((*matrix)[0])
	for col := 0; col < cols; col++ {
		availableSlot := -1
		for row := rows - 1; row > -1; row-- {
			currChar := (*matrix)[row][col]

			switch currChar {
			case EMPTY:
				if availableSlot == -1 {
					availableSlot = row
				}
			case ROUNDED_ROCK:
				if availableSlot != -1 {
					(*matrix)[row][col] = EMPTY
					(*matrix)[availableSlot][col] = ROUNDED_ROCK
					availableSlot--
					if (*matrix)[availableSlot][col] == CUBE_ROCK {
						availableSlot = -1
					}
				}
			case CUBE_ROCK:
				availableSlot = -1
			}
		}
	}
}

func rollEast(matrix *[][]rune) {
	rows, cols := len(*matrix), len((*matrix)[0])

	for row := 0; row < rows; row++ {
		availableSlot := -1
		for col := cols - 1; col > -1; col-- {
			currChar := (*matrix)[row][col]

			switch currChar {
			case EMPTY:
				if availableSlot == -1 {
					availableSlot = col
				}
			case ROUNDED_ROCK:
				if availableSlot != -1 {
					(*matrix)[row][col] = EMPTY
					(*matrix)[row][availableSlot] = ROUNDED_ROCK
					availableSlot--
					if (*matrix)[row][availableSlot] == CUBE_ROCK {
						availableSlot = -1
					}
				}
			case CUBE_ROCK:
				availableSlot = -1
			}
		}
	}
}

func prettyPrint(matrix [][]rune) {
	_, cols := len(matrix), len(matrix[0])
	fmt.Println(strings.Repeat("_", cols*2))
	for _, row := range matrix {
		for _, col := range row {
			fmt.Print(string(col), " ")
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("_", cols*2))
}

func loadInput(src string) ([][]rune, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, b := range line {
			row = append(row, b)
		}
		out = append(out, row)
	}

	return out, nil
}
