package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str, err := readInput()
	if err != nil {
		panic(err)
	}

	dp := make(map[string]int)

	fmt.Println(h(str, 0, "", &dp))
}

func h(str string, ind int, curr string, dp *map[string]int) int {
	key := fmt.Sprintf("%d,%s", ind, curr)

	if val, ok := (*dp)[key]; ok {
		return val
	}

	if curr == "QAQ" {
		return 1
	}

	if ind >= len(str) {
		return 0
	}

	var opt1, opt2 int

	if str[ind] == 'Q' && (curr == "" || curr == "QA") {
		opt1 = h(str, ind+1, fmt.Sprintf("%s%s", curr, "Q"), dp)
	} else if str[ind] == 'A' && curr == "Q" {
		opt1 = h(str, ind+1, "QA", dp)
	}

	opt2 = h(str, ind+1, curr, dp)

	(*dp)[key] = opt1 + opt2

	return opt1 + opt2
}

func readInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text(), nil
}
