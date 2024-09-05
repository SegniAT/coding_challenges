package main

import (
	"fmt"
	"slices"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Scan(&n)
		fmt.Scan(&k)

		dist := make([]int, k)
		for i := 0; i < k; i++ {
			fmt.Scan(&dist[i])
		}

		fmt.Println(maxMiceInHole(n, k, dist))
	}
}

// simulation
// not working
func maxMiceInHole(n, k int, dist []int) int {
	d := make([]int, k)
	copy(d, dist)
	slices.Sort(d)
	dist = d

	cnt, rem := 0, n
	for i := 0; rem > 0 && i < len(dist); i++ {
		rem -= dist[i]
		if rem > 0 {
			cnt++
		}
	}

	return cnt
}
