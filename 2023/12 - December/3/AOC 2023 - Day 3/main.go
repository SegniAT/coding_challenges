package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := getInput("input.txt")
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

	potentialGears := make(map[string][]int)

	sum := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if isNumber(matrix[row][col]) {
				leftCol := col - 1
				completeNum := ""
				gears := []string{}
				for col < cols && isNumber(matrix[row][col]) {
					completeNum = fmt.Sprintf("%s%c", completeNum, matrix[row][col])

					// check top
					if row-1 > -1 && isGear(matrix[row-1][col]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row-1, col))
					}

					// check bottom
					if row+1 < rows && isGear(matrix[row+1][col]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row+1, col))
					}
					col++
				}

				// check left side
				if leftCol > -1 {
					// top left
					if row-1 > -1 && isGear(matrix[row-1][leftCol]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row-1, leftCol))
					}
					// left
					if isGear(matrix[row][leftCol]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row, leftCol))
					}
					// bottom left
					if row+1 < rows && isGear(matrix[row+1][leftCol]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row+1, leftCol))
					}
				}

				// check right side
				if col < cols {
					// top right
					if row-1 > -1 && isGear(matrix[row-1][col]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row-1, col))
					}
					// right
					if isGear(matrix[row][col]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row, col))
					}
					// bottom right
					if row+1 < rows && isGear(matrix[row+1][col]) {
						gears = append(gears, fmt.Sprintf("%d,%d", row+1, col))
					}
				}

				completeNumInt, err := strconv.Atoi(completeNum)
				if err != nil {
					panic(err)
				}

				for _, gear := range gears {
					if nums, ok := potentialGears[gear]; ok {
						potentialGears[gear] = append(nums, completeNumInt)
					} else {
						potentialGears[gear] = []int{completeNumInt}
					}
				}
			}
		}
	}

	for _, nums := range potentialGears {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}

	fmt.Println("SUM: ", sum)

}

func isGear(char rune) bool {
	return char == '*'
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
