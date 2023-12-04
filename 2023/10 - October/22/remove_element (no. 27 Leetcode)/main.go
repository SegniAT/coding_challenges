package main

import "fmt"

func main() {
	nums := []int{1}
	val := 1
	res := removeElement(nums, val)
	fmt.Println(res)
}

func removeElement(nums []int, val int) int {
	res, left, right := 0, 0, len(nums)-1

	for left <= right {
		if nums[right] == val {
			right--
			continue
		} else if nums[left] == val {
			nums[left] = nums[right]
			nums[right] = val
		}
		left++
	}

	res = right + 1

	fmt.Println(nums)
	return res
}
