package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n)
	fmt.Scan(&k)

	fmt.Println(solve(n, k))
}

func solve(n, k int) int {
	minQ, maxQ := 1, n

	maxQuestions := 0
	for minQ <= maxQ {
		mid := minQ + (maxQ-minQ)/2

		// time taken to do questions from 1-mid...(1+2...mid)*5
		timeTaken := (mid * (mid + 1) / 2) * 5
		if timeTaken+k <= 4*60 {
			maxQuestions = max(maxQuestions, mid)
			minQ = mid + 1
		} else {
			maxQ = mid - 1
		}
	}

	return maxQuestions
}
