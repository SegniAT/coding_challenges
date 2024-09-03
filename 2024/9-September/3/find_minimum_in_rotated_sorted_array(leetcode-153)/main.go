package main

func main() {
}

/**
NOTE: if the middle num is less than the right most number
	-> it means it's increasing as it should...so look to the left including the middle element

NOTE: if the middle is greater than the right most number
	-> it means it's not normal...both the largest and smallest nums are to the right, don't include the middle num
**/

func findMin(nums []int) int {
	N := len(nums)
	left, right := 0, N-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
