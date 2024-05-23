package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := loadInput()

	for _, inp := range input {
		solve(inp)
	}
}

func solve(num int) {
	left, right := 0, num*3-1
	count := num / 2

	if !isEven(num) {
		count++
	}

	fmt.Println(count)
	for right > left {
		fmt.Printf("%d %d\n", left+1, right+1)
		left += 3
		right -= 3
	}
}

func isEven(num int) bool {
	return num%2 == 0
}

func loadInput() []int {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()
	tests, _ := strconv.Atoi(line)

	var out []int

	for range tests {
		scanner.Scan()
		line = scanner.Text()
		num, _ := strconv.Atoi(line)
		out = append(out, num)
	}
	return out
}
