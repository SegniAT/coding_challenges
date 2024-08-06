package main

func main() {
}

// Knapsack but added a way to avoid repitition since [2,1,1] and [1,2,1] should be counted as 1 combination
func change(amount int, coins []int) int {
	N := len(coins)

	cache := make([][]int, amount+1)
	for row := range cache {
		cache[row] = make([]int, N+1)
		for col := range cache[row] {
			cache[row][col] = -1
		}
	}

	var helper func(int, int) int

	helper = func(amountLeft, lastInd int) int {
		if amountLeft < 0 {
			return 0
		}

		if val := cache[amountLeft][lastInd]; val != -1 {
			return val
		}

		if amountLeft == 0 {
			return 1
		}

		ways := 0
		for i := lastInd; i < N; i++ {
			coin := coins[i]
			ways += helper(amountLeft-coin, i)
		}

		cache[amountLeft][lastInd] = ways

		return ways
	}

	return helper(amount, 0)
}
