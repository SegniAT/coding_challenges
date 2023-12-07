package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	times, err := parseInts(input[0])
	if err != nil {
		panic(err)
	}

	distances, err := parseInts(input[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("Time: ", times)
	fmt.Println("Distance: ", distances)

	races := len(distances)
	ans := 1

	for race := 0; race < races; race++ {
		recordsBroken := 0
		speed, time := 0, times[race]

		for ; speed < times[race]; speed++ {
			distance := speed * time
			if distance > distances[race] {
				recordsBroken++
			}
			time--
		}

		ans *= recordsBroken
	}

	fmt.Println("Ans: ", ans)
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
