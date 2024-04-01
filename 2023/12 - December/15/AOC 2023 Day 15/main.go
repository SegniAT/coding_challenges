package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Input: ", input)

	// sum := 0

	// for _, inp := range input {
	// 	sum += getHash(inp)
	// }

	// fmt.Println("SUM: ", sum)

	fmt.Println(getHash("rn=1"))

	// PART I
	// Attempt 1: 513158 (DONE!)

	// PART II
}

func getHash(str string) int {
	value := 0

	for _, char := range str {
		value += int(char)
		value *= 17
		value %= 256
	}

	return value
}

func loadInput(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	out := []string{}

	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")

	for _, l := range line {
		out = append(out, strings.TrimSpace(l))
	}

	return out, nil
}
