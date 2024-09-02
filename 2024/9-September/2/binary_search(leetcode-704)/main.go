package main

func main() {
}

func search(nums []int, target int) int {
	N := len(nums)
	left, right := 0, N-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[right] == target {
		return right
	}

	return -1
}
