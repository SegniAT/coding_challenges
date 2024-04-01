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
		// if ind == 1 {
		// 	continue
		// }
		ans += bruh(inp)
	}

	fmt.Println("ANS: ", ans)
	// PART I
	// First Attempt: 41859 (😭 thank GOD!)

	// PART II
	// First Attempt: 28299 (too low)
	// Second Attempt: 28357 (too low)

}

func bruh(pattern []string) int {
	rows, cols := len(pattern), len(pattern[0])

	ans := 0

	globalSmudgeInd := []int{}
	// check horizontal
	topInd := -1
	for i := 1; i < rows; i++ {
		top, bottom := i-1, i
		smudgeFound, smudgeInd := false, []int{}

		for top > -1 && bottom < rows {
			difference, firstDiffInd := diff(pattern[top], pattern[bottom])

			if !smudgeFound && difference == 1 {
				smudgeFound = true
				smudgeInd = []int{top, firstDiffInd}
			} else if difference > 0 {
				break
			}

			top--
			bottom++
		}

		if top < 0 || bottom >= rows && smudgeFound {
			topInd = i - 1
			globalSmudgeInd = smudgeInd
			break
		}
	}

	if topInd != -1 {
		ans += (topInd + 1) * 100
	}

	// fix smudge?
	if len(globalSmudgeInd) > 0 {
		row, col := globalSmudgeInd[0], globalSmudgeInd[1]

		str := pattern[row]

		newChar := '-'
		if str[col] == '.' {
			newChar = '#'
		} else {
			newChar = '.'
		}

		newStr := fmt.Sprintf("%s%c%s", str[:col], newChar, str[col+1:])

		pattern[row] = newStr
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

		smudgeFound, smudgeInd := false, []int{}
		if len(globalSmudgeInd) > 0 {
			smudgeFound = true
		}

		for left > -1 && right < cols {
			difference, firstDiffInd := diff(colsList[left], colsList[right])

			if !smudgeFound && difference == 1 {
				smudgeFound = true
				smudgeInd = []int{firstDiffInd, left}
			} else if difference > 0 {
				break
			}
			left--
			right++
		}

		if left < 0 || right >= cols && smudgeFound {
			leftInd = i - 1
			globalSmudgeInd = smudgeInd
			break
		}
	}

	if leftInd != -1 {
		ans += leftInd + 1
	}
	fmt.Println("globalshit: ", globalSmudgeInd)

	return ans
}

func diff(str1, str2 string) (int, int) {
	difference := 0
	firstDiffInd := -1
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if firstDiffInd == -1 {
				firstDiffInd = i
			}
			difference++
		}
	}

	return difference, firstDiffInd
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
