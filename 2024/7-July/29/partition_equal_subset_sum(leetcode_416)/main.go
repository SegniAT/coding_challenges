package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
	nums = []int{1}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	N := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	half := sum / 2

	cache := make([][]int, N)
	for i := range cache {
		cache[i] = make([]int, sum+1)
	}

	var helper func(int, int) bool
	helper = func(ind int, currSum int) bool {
		if ind < 0 || currSum < 0 {
			return false
		}

		if val := cache[ind][currSum]; val != 0 {
			return val == 1
		}

		if currSum == 0 {
			return true
		}

		opt1 := helper(ind-1, currSum-nums[ind])
		opt2 := helper(ind-1, currSum)

		if opt1 || opt2 {
			cache[ind][currSum] = 1
		} else {
			cache[ind][currSum] = 2
		}

		return opt1 || opt2
	}

	return helper(N-1, half)

}
