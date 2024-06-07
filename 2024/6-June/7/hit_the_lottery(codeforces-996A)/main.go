package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := loadInput()
	billTypes := []int{100, 20, 10, 5, 1}
	billCount := 0
	billInd := 0

	for n > 0 {
		if billTypes[billInd] > n {
			billInd++
		} else {
			n -= billTypes[billInd]
			billCount++
		}
	}

	fmt.Println(billCount)
}

func loadInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	num, _ := strconv.Atoi(line)
	return num
}
