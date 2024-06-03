package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	ans := numSubarrayProductLessThanK(nums, k)

	fmt.Println("ans: ", ans)
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	product := 1
	res := 0
	leftPtr := 0

	for rightPtr, num := range nums {
		product *= num

		for product >= k {
			product /= nums[leftPtr]
			leftPtr++
		}

		res += rightPtr - leftPtr + 1
	}

	return res
}
