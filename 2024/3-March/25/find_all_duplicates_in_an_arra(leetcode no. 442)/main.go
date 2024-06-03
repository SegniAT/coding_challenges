package main

import "math"

func main() {

}

// O(n) time & O(1) space
func findDuplicatesBetter(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		if num < 0 {
			num = int(math.Abs(float64(num)))
		}
		if nums[num-1] < 0 {
			res = append(res, num)
		} else {
			nums[num-1] *= -1
		}
	}
	return res
}

const (
	SEEN = 1
	DUP  = 2
)

func findDuplicates(nums []int) []int {
	hash := make(map[int]int)
	res := []int{}

	for _, num := range nums {
		if state := hash[num]; state == 0 {
			hash[num] = SEEN
		} else {
			hash[num] = DUP
		}
	}

	for num, state := range hash {
		if state == DUP {
			res = append(res, num)
		}
	}

	return res
}
