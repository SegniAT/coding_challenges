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

	for i := 2; i < len(input); i++ {
		line := input[i]
		equalsInd, openingParenthesisInd := strings.Index(line, "="), strings.Index(line, "(")
		commaInd, closingParenthesisInd := strings.Index(line, ","), strings.Index(line, ")")

		node := line[0 : equalsInd-1]
		dirs := []string{line[openingParenthesisInd+1 : commaInd], line[commaInd+2 : closingParenthesisInd]}

		nodeDirs[node] = dirs
	}

	startingNodes := []string{}
	for node := range nodeDirs {
		if node[2] == 'A' {
			startingNodes = append(startingNodes, node)
		}
	}

	steps := make([]int, len(startingNodes))

	for i := 0; i < len(startingNodes); i++ {
		steps[i] = findSteps(dirs, startingNodes[i], nodeDirs)
	}

	fmt.Println("Starting nodes: ", startingNodes)
	fmt.Println("Steps: ", steps)
	fmt.Println("LCM: ", lcmMultiple(steps))
	// 10668805667831

}

func findSteps(dirs, startNode string, nodeDirs map[string][]string) int {
	currNode := startNode
	i := 0
	steps := 0
	for true {
		if dirs[i] == LEFT {
			currNode = nodeDirs[currNode][0]
		}

		if dirs[i] == RIGHT {
			currNode = nodeDirs[currNode][1]
		}

		steps++

		if currNode[2] == 'Z' {
			break
		}

		i++
		if i >= len(dirs) {
			i = 0
		}
	}

	return steps
}

// Function to calculate the GCD (Greatest Common Divisor)
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to calculate the LCM (Least Common Multiple) of multiple numbers
func lcmMultiple(numbers []int) int {
	lcm := numbers[0]

	for i := 1; i < len(numbers); i++ {
		lcm = (lcm * numbers[i]) / gcd(lcm, numbers[i])
	}

	return lcm
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
