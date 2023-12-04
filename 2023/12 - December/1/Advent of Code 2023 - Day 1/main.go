package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// input := []string{"zoneight234"}
	// input, err := spelledToNum(input)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(input)

	input, err := getInputLines()
	if err != nil {
		panic(err)
	}

	input, err = spelledToNum(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(input)

	nums := []int{}

	for _, line := range input {
		firstDigit, lastDigit := 'n', 'n'

		for _, char := range line {
			if isNumber(char) {
				if firstDigit == 'n' {
					firstDigit = char
				}
				lastDigit = char
			}
		}

		num, err := strconv.Atoi(fmt.Sprintf("%c%c", firstDigit, lastDigit))
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("SUM: ", sum)

}

// part II
// eightwo -> 82 not 8wo
func spelledToNum(lines []string) ([]string, error) {
	spelled := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	const maxSpelledLen = 5

	res := make([]string, len(lines))

	for ind, line := range lines {

		for sp, num := range spelled {
			line = strings.ReplaceAll(line, sp, fmt.Sprintf("%c%d%c", sp[0], num, sp[len(sp)-1]))
		}

		res[ind] = line
	}

	return res, nil
}

func isNumber(char rune) bool {
	return '0' <= char && char <= '9'
}

func getInputLines() ([]string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	res := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return res, nil
}
