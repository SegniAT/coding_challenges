package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	GALAXY = '#'
	SPACE  = '.'
)

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Input: ")
	prettyPrint(input)

	fmt.Println()

	expanded := expandSpace(input)
	prettyPrint(expanded)

	rows, cols := len(expanded), len(expanded[0])
	galaxiesCoord := make(map[string][]int)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if expanded[row][col] == GALAXY {
				galaxiesCoord[fmt.Sprintf("%d", len(galaxiesCoord)+1)] = []int{row, col}
			}
		}
	}

	galaxies := make([]string, len(galaxiesCoord))
	for i := 0; i < len(galaxiesCoord); i++ {
		galaxies[i] = fmt.Sprintf("%d", i+1)
	}
	fmt.Println(galaxiesCoord)

	sumOfLengths := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxyOne, galaxyTwo := galaxiesCoord[fmt.Sprintf("%d", i+1)], galaxiesCoord[fmt.Sprintf("%d", j+1)]
			xDiff := int(math.Abs(float64(galaxyOne[0]) - float64(galaxyTwo[0])))
			yDiff := int(math.Abs(float64(galaxyOne[1]) - float64(galaxyTwo[1])))

			sumOfLengths += xDiff + yDiff
		}
	}

	fmt.Println("SUM OF LENGTHS: ", sumOfLengths)
}

func expandSpace(matrix [][]rune) (out [][]rune) {
	rows, cols := len(matrix), len(matrix[0])
	galaxyRows, galaxyCols := make([]bool, rows), make([]bool, cols)

	// mark galaxy rows and columns
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == GALAXY {
				galaxyRows[row] = true
				galaxyCols[col] = true
			}
		}
	}

	// expand columns and rows
	for row := 0; row < rows; row++ {
		newRow := []rune{}
		for col := 0; col < cols; col++ {
			newRow = append(newRow, matrix[row][col])
			if !galaxyCols[col] {
				newRow = append(newRow, matrix[row][col])
			}
		}
		out = append(out, newRow)
		if !galaxyRows[row] {
			out = append(out, newRow)
		}
	}

	return out
}

func prettyPrint(matrix [][]rune) {
	line := strings.Repeat("_", len(matrix)*2)
	fmt.Println(line)
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Println()
	}
	fmt.Println(line)
}

func loadInput(src string) ([][]rune, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := [][]rune{}

	for scanner.Scan() {
		row := scanner.Bytes()
		runeRow := []rune{}
		for _, r := range row {
			runeRow = append(runeRow, rune(r))
		}
		out = append(out, runeRow)
	}

	return out, nil
}
