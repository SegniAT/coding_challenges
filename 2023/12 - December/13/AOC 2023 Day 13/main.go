package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	ans := 0
	for _, inp := range input {
		ans += bruh(inp)
	}

	fmt.Println("ANS: ", ans)
	// PART I
	// First Attempt: 41859 (😭 thank GOD!)

}

func bruh(pattern []string) int {
	rows, cols := len(pattern), len(pattern[0])

	ans := 0

	// check horizontal
	topInd := -1
	for i := 1; i < rows; i++ {
		top, bottom := i-1, i

		for top > -1 && bottom < rows && pattern[top] == pattern[bottom] {
			top--
			bottom++
		}

		if top < 0 || bottom >= rows {
			topInd = i - 1
			break
		}
	}

	if topInd != -1 {
		ans += (topInd + 1) * 100
	}

	// check vertical

	// build full column images
	colsList := []string{}
	for col := 0; col < cols; col++ {
		colStr := ""
		for row := 0; row < rows; row++ {
			colStr = fmt.Sprintf("%s%c", colStr, pattern[row][col])
		}
		colsList = append(colsList, colStr)
	}

	// same steps as for the horizontal check
	leftInd := -1
	for i := 1; i < cols; i++ {
		left, right := i-1, i

		for left > -1 && right < cols && colsList[left] == colsList[right] {
			left--
			right++
		}

		if left < 0 || right >= cols {
			leftInd = i - 1
			break
		}
	}

	if leftInd != -1 {
		ans += leftInd + 1
	}

	return ans
}

func loadInput(src string) ([][]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	patterns := [][]string{}

	patternLines := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			patterns = append(patterns, patternLines)
			patternLines = []string{}
			continue
		}

		patternLines = append(patternLines, line)
	}

	if len(patternLines) > 0 {
		patterns = append(patterns, patternLines)
	}

	return patterns, nil
}
