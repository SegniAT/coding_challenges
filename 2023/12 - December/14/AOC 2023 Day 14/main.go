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
	fmt.Println("Input: ")
	prettyPrint(input)
	rows, cols := len(input), len(input[0])
	totalLoad := 0

	for col := 0; col < cols; col++ {
		availableSlot := -1
		for row := 0; row < rows; row++ {
			currChar := input[row][col]

			switch currChar {
			case EMPTY:
				if availableSlot == -1 {
					availableSlot = row
				}
			case ROUNDED_ROCK:
				if availableSlot != -1 {
					input[row][col] = EMPTY
					input[availableSlot][col] = ROUNDED_ROCK
					totalLoad += rows - availableSlot
					availableSlot++
					if input[availableSlot][col] == CUBE_ROCK {
						availableSlot = -1
					}
				} else {
					totalLoad += rows - row
				}

			case CUBE_ROCK:
				availableSlot = -1
			}
		}
	}

	fmt.Println()
	fmt.Println("Modified: ")
	prettyPrint(input)

	fmt.Println("Total load: ", totalLoad)
	// Part I
	// Attempt 1: 109466 (DONE!!)
}

func prettyPrint(matrix [][]rune) {
	_, cols := len(matrix), len(matrix[0])
	fmt.Println(strings.Repeat("_", cols))
	for _, row := range matrix {
		for _, col := range row {
			fmt.Print(string(col), " ")
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("_", cols))
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
