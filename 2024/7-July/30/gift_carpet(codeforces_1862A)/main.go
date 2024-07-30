package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NAME = "vika"

func main() {
	inputs, err := readInput()
	if err != nil {
		fmt.Println("error reading the input")
		panic(err)
	}

	for _, input := range inputs {
		N, M := len(input), len(input[0])
		if M < len(NAME) {
			fmt.Println("NO")
			continue
		}

		targetInd := 0
		found := false

		for col := 0; col < M && !found; col++ {
			for row := 0; row < N && !found; row++ {
				if input[row][col] == NAME[targetInd] {
					targetInd++
					if targetInd >= len(NAME) {
						found = true
					}
					break
				}
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func readInput() ([][]string, error) {
	out := [][]string{}
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	tests, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return out, err
	}

	for tests > 0 {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		N, err := strconv.Atoi(line[0])
		if err != nil {
			return out, err
		}

		rows := []string{}
		for N > 0 {
			scanner.Scan()
			rows = append(rows, scanner.Text())
			N--
		}

		out = append(out, rows)

		tests--
	}

	return out, nil
}
