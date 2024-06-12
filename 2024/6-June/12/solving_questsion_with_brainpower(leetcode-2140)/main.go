package main

import "fmt"

func main() {
	inp := [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}
	fmt.Println(mostPoints(inp))
	inp = [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
	fmt.Println(mostPoints(inp))
}

// TOP DOWN
/*
func mostPoints(questions [][]int) int64 {
	dp := make(map[int]int64)
	return helper(questions, 0, &dp)
}

func helper(questions [][]int, ind int, dp *map[int]int64) int64 {
	if ind >= len(questions) {
		return 0
	}

	if val, ok := (*dp)[ind]; ok {
		return val
	}

	points, brainpower := questions[ind][0], questions[ind][1]

	take := helper(questions, ind+brainpower+1, dp) + int64(points)
	skip := helper(questions, ind+1, dp)

	(*dp)[ind] = max(take, skip)
	return (*dp)[ind]
}


*/

func mostPoints(questions [][]int) int64 {
	size := len(questions)
	dp := make([]int64, len(questions))
	for i := range size {
		dp[i] = int64(questions[i][0])
	}

	var brainPower, nextInd int
	for i := range size {
		brainPower = questions[i][1]
		nextInd = i + brainPower + 1
		if nextInd < size {
			dp[nextInd] = max(dp[nextInd], dp[i]+int64(questions[nextInd][0]))
		}
	}

	var mostPts int64
	for i := range size {
		mostPts = max(mostPts, dp[i])
	}
	return mostPts
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
