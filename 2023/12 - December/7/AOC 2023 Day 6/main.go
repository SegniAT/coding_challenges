package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput("sample_input.txt")
	if err != nil {
		panic(err)
	}

	maxTime, err := parseIntsPartTwo(input[0])
	if err != nil {
		panic(err)
	}

	recordDistance, err := parseIntsPartTwo(input[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("Time: ", maxTime)
	fmt.Println("Distance: ", recordDistance)

	ans := 1

	ans *= waysToWinBetter(maxTime, recordDistance)

	fmt.Println("Ans: ", ans)
}

func waysToWin(maxTime, recordDistance int) int {
	recordsBroken := 0
	speed, time := 0, maxTime

	for ; speed < maxTime; speed++ {
		distance := speed * time
		if distance > recordDistance {
			recordsBroken++
		}
		time--
	}
	return recordsBroken
}

func waysToWinBetter(maxTime, recordDistance int) int {
	firstWinTime := 0
	speed, time := 0, maxTime

	for ; speed < maxTime; speed++ {
		distance := speed * time
		if distance > recordDistance {
			firstWinTime = speed
			break
		}
		time--
	}

	// + 1 is necessary for some reason
	return maxTime - firstWinTime*2 + 1
}

func parseInts(str string) ([]int, error) {
	out := []int{}
	currNum := ""

	colonInd := strings.Index(str, ":")
	str = str[colonInd:]
	str = fmt.Sprintf("%s ", str)

	for _, char := range str {
		if isNum(char) {
			currNum = fmt.Sprintf("%s%c", currNum, char)
		} else if currNum != "" {
			num, err := strconv.Atoi(currNum)
			if err != nil {
				panic(err)
			}
			out = append(out, num)
			currNum = ""
		}
	}
	return out, nil
}

func parseIntsPartTwo(str string) (int, error) {
	numStr := ""
	for _, char := range str {
		if isNum(char) {
			numStr = fmt.Sprintf("%s%c", numStr, char)
		}
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1, err
	}

	return num, nil
}

func isNum(char rune) bool {
	return '0' <= char && char <= '9'
}

func readInput(src string) ([]string, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	out := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, nil
}
