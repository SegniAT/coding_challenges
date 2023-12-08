package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	LEFT  = 'L'
	RIGHT = 'R'
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	dirs := input[0]
	nodeDirs := make(map[string][]string)

	fmt.Println(dirs)
	for i := 2; i < len(input); i++ {
		line := input[i]
		equalsInd, openingParenthesisInd := strings.Index(line, "="), strings.Index(line, "(")
		commaInd, closingParenthesisInd := strings.Index(line, ","), strings.Index(line, ")")

		node := line[0 : equalsInd-1]
		dirs := []string{line[openingParenthesisInd+1 : commaInd], line[commaInd+2 : closingParenthesisInd]}

		nodeDirs[node] = dirs
	}

	dirsCount := len(dirs)
	i := 0
	currNode := "AAA"
	steps := 0
	for true {
		dir := dirs[i]

		if dir == LEFT {
			currNode = nodeDirs[currNode][0]
		}

		if dir == RIGHT {
			currNode = nodeDirs[currNode][1]
		}

		steps++
		if currNode == "ZZZ" {
			break
		}
		i++
		if i >= dirsCount {
			i = 0
		}
	}

	fmt.Println("Steps: ", steps)

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
