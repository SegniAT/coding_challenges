package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, k := loadInput()
	//fmt.Println(n, k)

	// 1 + 2 + 3 ... pattern n(n+1)/2
	left, right := 1, n

	for right > left {
		mid := (right + left) / 2
		candyCount := (mid * (mid + 1) / 2) - (n - mid)

		if candyCount >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	//fmt.Println("right: ", right)

	candyCount := right * (right + 1) / 2

	candyEaten := candyCount - k
	fmt.Println(candyEaten)

}

func loadInput() (n, k int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := strings.Split(scanner.Text(), " ")

	n, _ = strconv.Atoi(line[0])
	k, _ = strconv.Atoi(line[1])

	return
}
