package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs, err := readInput()
	if err != nil {
		panic(err)
	}

	dp := make(map[string][]int)

	var helper func([]int) []int
	helper = func(abc []int) []int {
		if valid(abc) {
			return abc
		}

		a, b, c := abc[0], abc[1], abc[2]
		key := fmt.Sprintf("%d %d %d", a, b, c)
		if res, ok := dp[key]; ok {
			return res
		}

		res := []int{0, 0, 0}

		// remove 1,2
		if a > 0 && b > 0 {
			opt1 := helper([]int{a - 1, b - 1, c + 1})
			//fmt.Println("opt1: ", opt1)

			for i := 0; i < 3; i++ {
				if opt1[i] > 0 {
					res[i] = 1
				}
			}
		}

		// remove 1,3
		if a > 0 && c > 0 {
			opt2 := helper([]int{a - 1, b + 1, c - 1})
			//fmt.Println("opt2: ", opt2)

			for i := 0; i < 3; i++ {
				if opt2[i] > 0 {
					res[i] = 1
				}

			}
		}

		// remove 2,3
		if b > 0 && c > 0 {
			opt3 := helper([]int{a + 1, b - 1, c - 1})
			//fmt.Println("opt3: ", opt3)

			for i := 0; i < 3; i++ {
				if opt3[i] > 0 {
					res[i] = 1
				}
			}

		}

		dp[key] = res
		return res
	}

	for _, input := range inputs {
		res := helper(input)
		fmt.Println(res[0], res[1], res[2])
	}

	//fmt.Println(dp)
}

func valid(abc []int) bool {
	zeroes := 0
	for _, num := range abc {
		if num == 0 {
			zeroes++
		}
	}

	return zeroes == 2
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
		split := strings.Split(scanner.Text(), " ")

		one, err := strconv.Atoi(split[0])
		if err != nil {
			return out, err
		}
		two, err := strconv.Atoi(split[1])
		if err != nil {
			return out, err
		}
		three, err := strconv.Atoi(split[2])
		if err != nil {
			return out, err
		}

		out = append(out, []int{one, two, three})
	}

	return out, nil
}
