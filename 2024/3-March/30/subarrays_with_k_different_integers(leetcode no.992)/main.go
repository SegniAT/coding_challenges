package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 2, 3}
	k := 2

	fmt.Println(subarraysWithKDistinct(nums, k))

	nums = []int{1, 2, 1, 3, 4}
	k = 3

	fmt.Println(subarraysWithKDistinct(nums, k))

	nums = []int{1, 2}
	k = 1

	fmt.Println(subarraysWithKDistinct(nums, k))
}

func subarraysWithKDistinct(nums []int, k int) int {
	count := make(map[int]int)
	leftPtr := 0
	subArrays := 0

	for _, num := range nums {
		count[num]++

		for len(count) > k {
			leftNum := nums[leftPtr]
			count[leftNum]--
			if count[leftNum] == 0 {
				delete(count, leftNum)
			}
			leftPtr++
		}

		flag := len(count) == k
		oldLeftPtr := leftPtr

		for len(count) == k {
			leftNum := nums[leftPtr]
			count[leftNum]--
			if count[leftNum] == 0 {
				delete(count, leftNum)
			}
			leftPtr++
		}

		if flag {
			subArrays += (leftPtr - oldLeftPtr)
			for i := leftPtr - 1; i >= oldLeftPtr; i-- {
				count[nums[i]]++
			}
			leftPtr = oldLeftPtr
		}

	}

	return subArrays
}
