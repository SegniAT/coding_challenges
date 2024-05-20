package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	tests, err := loadInput()
	if err != nil {
		log.Fatal("error reading input: ", err)
	}

outer:
	for _, test := range tests {
		rows := len(test)
		count := make(map[rune]int)

		for _, row := range test {
			for _, char := range row {
				count[char]++
			}
		}

		for _, val := range count {
			if val%rows != 0 {
				fmt.Println("NO")
				continue outer
			}
		}

		fmt.Println("YES")
	}

}

func loadInput() ([][]string, error) {
	out := [][]string{}

	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	tests, err := strconv.Atoi(input.Text())
	if err != nil {
		return out, err
	}

	for range tests {
		input.Scan()
		rows, err := strconv.Atoi(input.Text())
		if err != nil {
			return out, err
		}

		test := []string{}

		for range rows {
			input.Scan()
			test = append(test, input.Text())
		}

		out = append(out, test)
	}

	return out, nil
}
