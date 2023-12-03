package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := getInput("sample_input.txt")
	if err != nil {
		panic(err)
	}

	rows, cols := len(input), len(input[0])
	matrix := make([][]rune, rows)

	for rowInd, row := range input {
		matrix[rowInd] = make([]rune, cols)
		for colInd, col := range row {
			matrix[rowInd][colInd] = col
		}
	}

	sum := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if isNumber(matrix[row][col]) {
				leftCol := col - 1
				completeNum := ""
				hasAdjacentSymbol := false
				for col < cols && isNumber(matrix[row][col]) {
					completeNum = fmt.Sprintf("%s%c", completeNum, matrix[row][col])

					if !hasAdjacentSymbol {
						// check top
						if row-1 > -1 && isSymbol(matrix[row-1][col]) {
							hasAdjacentSymbol = true
						}

						// check bottom
						if row+1 < rows && isSymbol(matrix[row+1][col]) {
							hasAdjacentSymbol = true
						}
					}
					col++
				}

				// check left side
				if !hasAdjacentSymbol && leftCol > -1 {
					// top left
					if row-1 > -1 && isSymbol(matrix[row-1][leftCol]) {
						hasAdjacentSymbol = true
					}
					// left
					if isSymbol(matrix[row][leftCol]) {
						hasAdjacentSymbol = true
					}
					// bottom left
					if row+1 < rows && isSymbol(matrix[row+1][leftCol]) {
						hasAdjacentSymbol = true
					}
				}

				// check right side
				if !hasAdjacentSymbol && col < cols {
					// top right
					if row-1 > -1 && isSymbol(matrix[row-1][col]) {
						hasAdjacentSymbol = true
					}
					// right
					if isSymbol(matrix[row][col]) {
						hasAdjacentSymbol = true
					}
					// bottom right
					if row+1 < rows && isSymbol(matrix[row+1][col]) {
						hasAdjacentSymbol = true
					}

				}

				if hasAdjacentSymbol {
					// fmt.Println(completeNum)
					completeNumInt, err := strconv.Atoi(completeNum)
					if err != nil {
						panic(err)
					}

					sum += completeNumInt
				}
			}
		}
	}

	fmt.Println("SUM: ", sum)

}

func isSymbol(char rune) bool {
	return !isNumber(char) && char != '.'
}

func isNumber(char rune) bool {
	return '0' <= char && char <= '9'
}

func getInput(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}

	return out, nil
}
