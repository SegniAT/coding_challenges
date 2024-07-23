package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tests, err := readInput()
	if err != nil {
		panic(err)
	}

	for _, candies := range tests {
		n := len(candies)
		dp := make(map[[2]int]bool)

		totalWeight := 0
		for _, candy := range candies {
			totalWeight += candy
		}

		if totalWeight%2 != 0 {
			fmt.Println("NO")
			continue
		}

		res := canPartition(candies, n, 0, totalWeight/2, &dp)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func canPartition(candies []int, n, ind, target int, memo *map[[2]int]bool) bool {
	if target == 0 {
		return true
	}

	if ind >= n || target < 0 {
		return false
	}

	if val, ok := (*memo)[[2]int{ind, target}]; ok {
		return val
	}

	opt1 := canPartition(candies, n, ind+1, target-candies[ind], memo)
	opt2 := canPartition(candies, n, ind+1, target, memo)

	(*memo)[[2]int{ind, target}] = opt1 || opt2

	return opt1 || opt2
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
		ln, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return out, err
		}

		test := make([]int, ln)

		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")

		for i := range ln {
			test[i], err = strconv.Atoi(nums[i])
			if err != nil {
				return out, err
			}
		}

		out = append(out, test)
	}

	return out, nil
}
