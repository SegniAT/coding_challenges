package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs, err := readInput()
	if err != nil {
		log.Fatal("error reading input: ", err)
	}

	for _, input := range inputs {
		solve(input)
	}

}

func solve(input []int) {
	prefixSum := generatePrefixSum(input)

	ln := len(prefixSum)
	for i := 0; i < ln; i++ {
		left := prefixSum[i]
		right := prefixSum[ln-1] - left

		if left == right {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}

func generatePrefixSum(input []int) []int {
	ln := len(input)

	prefixSum := make([]int, ln)
	prefixSum[0] = input[0]

	for i := 1; i < ln; i++ {
		prefixSum[i] = input[i] + prefixSum[i-1]
	}
	return prefixSum
}

func readInput() ([][]int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	out := [][]int{}

	scanner.Scan()
	tests, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return out, err
	}

	for range tests {
		scanner.Scan()
		scanner.Text()

		scanner.Scan()
		numsLine := strings.Split(scanner.Text(), " ")
		nums, err := strArrToIntArr(numsLine)
		if err != nil {
			return out, err
		}
		out = append(out, nums)
	}

	return out, nil
}

func strArrToIntArr(arr []string) ([]int, error) {
	out := []int{}

	for _, str := range arr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return out, err
		}
		out = append(out, num)

	}

	return out, nil
}
