package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	cache := make(map[int]int)

	for _, num := range nums {
		cache[num] = -1
	}

	var helper func(int) int
	helper = func(num int) int {
		if _, ok := cache[num]; !ok {
			return 0
		}

		if val := cache[num]; val != -1 {
			return val
		}

		cache[num] = 1 + helper(num+1)

		return cache[num]
	}

	for _, num := range nums {
		helper(num)
	}

	longest := -1
	for _, val := range cache {
		if val > longest {
			longest = val
		}
	}

	return longest
}
