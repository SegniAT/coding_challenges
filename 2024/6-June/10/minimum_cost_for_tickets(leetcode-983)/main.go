package main

import (
	"fmt"
	"math"
)

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	fmt.Println(mincostTickets(days, costs))

	fmt.Println()
	days = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs = []int{2, 7, 15}
	fmt.Println(mincostTickets(days, costs))
}

var indexDaysMap = []int{1, 7, 30}

func mincostTickets(days []int, costs []int) int {
	dp := []int{}
	for range len(days) {
		dp = append(dp, -1)
	}
	res, _ := helper(days, costs, 0, 0, &dp)
	fmt.Println(dp)
	return res
}
func helper(days []int, costs []int, ind int, cost int, dp *[]int) (int, bool) {
	if ind >= len(days) {
		return cost, true
	}

	if val := (*dp)[ind]; val != -1 {
		return val, false
	}

	minCost := math.MaxInt
	stepCost := math.MaxInt
	fromLeaf := false
	for i := 0; i < 3; i++ {
		nextMaxInd := findNextMax(days, ind, days[ind]+indexDaysMap[i])
		child, isFromLeaf := helper(days, costs, nextMaxInd, costs[i], dp)
		stepCost = costs[i]
		if days[ind] == 20 {
			fmt.Println("child: ", child)
			fmt.Println("stepCost: ", stepCost)
		}
		if child+stepCost < minCost {
			minCost = child + stepCost
			fromLeaf = isFromLeaf
		}
	}

	if fromLeaf {
		minCost /= 2
	}

	(*dp)[ind] = minCost

	return minCost, false
}

func findNextMax(days []int, ind int, day int) int {
	for ; ind < len(days) && days[ind] <= day; ind++ {
	}

	return ind
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
