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

	cards := len(input)
	totalPoints := 0

	copies := make([]int, cards+1)
	for ind := 1; ind < cards+1; ind++ {
		copies[ind] = 1
	}

	for cardInd, card := range input {
		cardInd = cardInd + 1
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

		// fmt.Printf("Card %d - %d\n", cardInd, winningNumsFromOurNums)

		last := cardInd + winningNumsFromOurNums
		for ind := cardInd + 1; ind < len(copies) && ind <= last; ind++ {
			copies[ind] += copies[cardInd]
		}

		if winningNumsFromOurNums > 0 {
			totalPoints += int(math.Pow(2, float64(winningNumsFromOurNums)-1))
		}
	}

	totalCopies := 0
	for _, copy := range copies {
		totalCopies += copy
	}

	fmt.Println("Copies: ", copies)
	fmt.Println("Total copies: ", totalCopies)
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
