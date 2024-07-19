package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs, err := readInput()
	if err != nil {
		fmt.Println("Error reading the input")
		panic(err)
	}

	for _, nums := range inputs {
		m := len(nums)
		dp := make([][]int, m)
		for row := range m {
			dp[row] = make([]int, m)
			for col := range m {
				dp[row][col] = -1
			}
		}

		minm, maxm := math.MaxInt, math.MinInt

		for _, num := range nums {
			minm = min(minm, num)
			maxm = max(maxm, num)
		}

		var helper func(int, int, int, int) int
		helper = func(l, r, step, found int) int {
			if found == 2 {
				return step
			}

			if res := dp[l][r]; res != -1 {
				return res
			}

			// option 1
			opt1Found := found
			if nums[l] == minm || nums[l] == maxm {
				opt1Found++
			}
			opt1 := helper(l+1, r, step+1, opt1Found)

			// option 2
			opt2Found := found
			if nums[r] == minm || nums[r] == maxm {
				opt2Found++
			}
			opt2 := helper(l, r-1, step+1, opt2Found)

			dp[l][r] = min(opt1, opt2)
			return dp[l][r]

		}

		ans := helper(0, m-1, 0, 0)
		fmt.Println(ans)
	}

}

/*
		func helper(nums []int, l, r, step, minm, maxm, found int) int {
	}
*/
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
