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

	// fmt.Println("Input: ")
	// prettyPrint(input)

	fmt.Println()

	expandedRowInd, expandedColInd := spaceExpandedIndexes(input, 1_000_000)

	rows, cols := len(input), len(input[0])
	galaxyIndexes := [][]int{}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if input[row][col] == GALAXY {
				galaxyIndexes = append(galaxyIndexes, []int{expandedRowInd[row], expandedColInd[col]})
			}
		}
	}

	// fmt.Println(galaxyIndexes)

	sumOfLengths := 0
	for i := 0; i < len(galaxyIndexes); i++ {
		for j := i + 1; j < len(galaxyIndexes); j++ {
			galaxyOne, galaxyTwo := galaxyIndexes[i], galaxyIndexes[j]
			sumOfLengths += getDistance(galaxyOne, galaxyTwo)
		}
	}

	fmt.Println("SUM OF LENGTHS: ", sumOfLengths)
}

func getDistance(coordOne, coordTwo []int) int {
	xDiff := int(math.Abs(float64(coordOne[0]) - float64(coordTwo[0])))
	yDiff := int(math.Abs(float64(coordOne[1]) - float64(coordTwo[1])))
	return xDiff + yDiff
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

func spaceExpandedIndexes(matrix [][]rune, expansionSize int) ([]int, []int) {
	rows, cols := len(matrix), len(matrix[0])
	rowGalaxyPrefixSum, colGalaxyPrefixSum := make([]int, rows), make([]int, cols)

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

	for row := 0; row < rows; row++ {
		ind := 0
		for col := 0; col < cols; col++ {
			colGalaxyPrefixSum[col] = ind
			if !galaxyCols[col] {
				ind += expansionSize
			} else {
				ind++
			}
		}
	}

	for col := 0; col < cols; col++ {
		ind := 0
		for row := 0; row < rows; row++ {
			rowGalaxyPrefixSum[row] = ind
			if !galaxyRows[row] {
				ind += expansionSize
			} else {
				ind++
			}
		}
	}

	return rowGalaxyPrefixSum, colGalaxyPrefixSum
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
