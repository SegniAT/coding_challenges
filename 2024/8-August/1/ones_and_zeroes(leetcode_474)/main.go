package main

import "fmt"

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3

	fmt.Println(findMaxForm(strs, m, n))

}

const (
	MAX_K = 600
	MAX_m = 100
	MAX_n = 100
)

func findMaxForm(strs []string, m int, n int) int {
	N := len(strs)
	counts := [][2]int{}
	for _, str := range strs {
		counts = append(counts, countOnesZeroes(str))
	}

	cache := make([][][]int, MAX_K+1)

	for i := range cache {
		cache[i] = make([][]int, MAX_m+1)
		for j := range cache[i] {
			cache[i][j] = make([]int, MAX_n+1)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}

	var helper func(int, int, int) int

	helper = func(k, zeroesLeft, onesLeft int) int {
		if k < 0 || (zeroesLeft == 0 && onesLeft == 0) {
			return 0
		}

		if val := cache[k][zeroesLeft][onesLeft]; val != -1 {
			return val
		}

		zeroes, ones := counts[k][0], counts[k][1]

		if zeroesLeft-zeroes < 0 || onesLeft-ones < 0 {
			return helper(k-1, zeroesLeft, onesLeft)
		}

		// take, don't take
		opt1 := helper(k-1, zeroesLeft-zeroes, onesLeft-ones) + 1
		opt2 := helper(k-1, zeroesLeft, onesLeft)

		cache[k][zeroesLeft][onesLeft] = max(opt1, opt2)

		return max(opt1, opt2)
	}

	return helper(N-1, m, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countOnesZeroes(str string) [2]int {
	zeroes, ones := 0, 0
	for i := range str {
		if str[i] == '0' {
			zeroes++
		} else if str[i] == '1' {
			ones++
		}
	}

	return [2]int{zeroes, ones}
}
