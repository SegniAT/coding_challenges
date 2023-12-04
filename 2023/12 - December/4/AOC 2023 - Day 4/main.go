package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := loadInput("input.txt")
	if err != nil {
		panic(err)
	}

	totalPoints := 0

	for _, card := range input {
		// fmt.Println("Card: ", cardInd+1)
		colonIndex := strings.Index(card, ":")
		card = strings.TrimSpace(card[colonIndex+1:])
		temp := strings.Split(card, "|")

		winningNumbers, err := parseNums(temp[0])
		if err != nil {
			panic(err)
		}
		numbersWeHave, err := parseNums(temp[1])
		if err != nil {
			panic(err)
		}

		// fmt.Println(winningNumbers, numbersWeHave)

		winningNumbersMap := make(map[int]int)

		for _, winningNumber := range winningNumbers {
			winningNumbersMap[winningNumber]++
		}

		winningNumsFromOurNums := 0

		for _, num := range numbersWeHave {
			if _, ok := winningNumbersMap[num]; ok {
				winningNumsFromOurNums++
			}
		}

		if winningNumsFromOurNums > 0 {
			totalPoints += int(math.Pow(2, float64(winningNumsFromOurNums)-1))
		}
	}

	fmt.Println("Total points: ", totalPoints)

}

func parseNums(str string) ([]int, error) {
	nums := []int{}
	currNum := ""

	for _, char := range str {
		if isNumber(char) {
			currNum = fmt.Sprintf("%s%c", currNum, char)
		} else if currNum != "" {
			currNumInt, err := strconv.Atoi(currNum)
			if err != nil {
				return nil, err
			}
			nums = append(nums, currNumInt)
			currNum = ""
		}
	}

	if currNum != "" {
		currNumInt, err := strconv.Atoi(currNum)
		if err != nil {
			return nil, err
		}
		nums = append(nums, currNumInt)
		currNum = ""
	}

	return nums, nil
}

func isNumber(char rune) bool {
	return '0' <= char && char <= '9'
}

func loadInput(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	out := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		out = append(out, line)
	}

	return out, nil
}
