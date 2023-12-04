package main

import "fmt"

func main() {
	nums := []int{1, 2}
	fmt.Println(removeDuplicatesFaster(nums))
}

func removeDuplicates(nums []int) int {
	slow, fast := 0, 1

	for fast < len(nums) {
		if nums[slow] < nums[fast] {
			nums[slow+1], nums[fast] = nums[fast], nums[slow+1]
			slow++
			fast = slow + 1
		} else {
			fast++
		}
	}

	fmt.Println(nums)

	return slow + 1
}

func removeDuplicatesFaster(nums []int) int {
	slow := 0

	for fast := range nums {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
