package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	coins, amount := []int{1, 2, 5}, 11
	fmt.Println(coinChange(coins, amount))
	coins, amount = []int{2}, 3
	fmt.Println(coinChange(coins, amount))
	coins, amount = []int{3, 7, 405, 436}, 8839
	fmt.Println(coinChange(coins, amount))
}

// DP - Knapsack
// has to be Int32 so that it doesn't not wrap to negative
// works fine for the res == math.MaxInt32 too since we take the minimums even if there were additions
func coinChange(coins []int, amount int) int {
	cache := make([]int, amount+1)
	for i := range cache {
		cache[i] = -1
	}

	N := len(coins)
	sort.Ints(coins)

	var helper func(int) int
	helper = func(amountLeft int) int {
		if val := cache[amountLeft]; val != -1 {
			return val
		}

		if amountLeft < 0 {
			return math.MaxInt32
		}

		if amountLeft == 0 {
			return 0
		}

		minCoins := math.MaxInt32
		for i := N - 1; i > -1; i-- {
			coin := coins[i]
			if coin > amountLeft {
				continue
			}
			minCoins = min(minCoins, helper(amountLeft-coin)+1)
		}

		cache[amountLeft] = minCoins

		return minCoins
	}

	res := helper(amount)
	if res == math.MaxInt32 {
		res = -1
	}

	return res
}

// GREEDY - not always right
/*
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	minCoinsUsed := 0

	for i := len(coins) - 1; i > -1; i-- {
		if amount <= 0 {
			break
		}

		if coins[i] <= amount {
			q := amount / coins[i]
			minCoinsUsed += q
			amount -= coins[i] * q
		}
	}

	if amount != 0 {
		return -1
	}

	return minCoinsUsed

}
*/
