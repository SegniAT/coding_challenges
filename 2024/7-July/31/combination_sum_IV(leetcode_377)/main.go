package main

func main() {
}

// TOP DOWN
func combinationSum4(nums []int, target int) int {
	cache := make(map[int]int)

	var helper func(int) int
	helper = func(targetLeft int) int {
		if targetLeft < 0 {
			return 0
		}

		if val, ok := cache[targetLeft]; ok {
			return val
		}
		if targetLeft == 0 {
			return 1
		}

		count := 0
		for _, num := range nums {
			count += helper(targetLeft - num)
		}

		cache[targetLeft] = count

		return count
	}

	return helper(target)
}
