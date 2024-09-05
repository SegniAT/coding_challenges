package main

func main() {
}

func search(nums []int, target int) int {
	N := len(nums)
	left, right := 0, N-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// if the left side is increasing
		if nums[left] <= nums[mid] {
			// if the target is in the left side (between indexes left and mid)
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right side is increasing
			// if the target is in the right side (between indexes mid and right)
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
