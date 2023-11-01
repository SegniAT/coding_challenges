package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	profit := 0.0

	for i := 0; i < len(prices)-1; i++ {
		profit += math.Max(0, float64(prices[i+1])-float64(prices[i]))
	}

	return int(profit)
}
