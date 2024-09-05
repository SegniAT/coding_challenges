package main

import (
	"fmt"
)

// TLE...don't care anymore
func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		// must use all
		var sum uint64

		for range n {
			var n uint64
			fmt.Scan(&n)
			sum += n
		}

		if isPerfectSquare(sum) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isPerfectSquare(num uint64) bool {
	if num < 0 {
		return false
	}
	var low, high uint64
	low, high = 0, 1e9

	for low <= high {
		mid := low + (high-low)/2

		prod := mid * mid

		if prod == num {
			return true
		} else if prod < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
