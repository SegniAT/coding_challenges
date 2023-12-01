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

		for left := 0; left < len(line); left++ {
			if isNumber(rune(line[left])) {
				continue
			}
			for right := left; right < len(line); right++ {
				sliced := line[left : right+1]
				if len(sliced) > maxSpelledLen {
					break
				}

				if num, ok := spelled[sliced]; ok {
					fmt.Println(sliced, right)
					line = strings.Replace(line, sliced, fmt.Sprintf("%d%c", num, rune(line[right])), 1)
					break
				}
			}
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
