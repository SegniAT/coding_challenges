package main

func main() {

}

func removeDuplicates(nums []int) int {
	left := 0

	for right := range nums {
		if left == 0 || left == 1 || nums[left] != nums[right] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
