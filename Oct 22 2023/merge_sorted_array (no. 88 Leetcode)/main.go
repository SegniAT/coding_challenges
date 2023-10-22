package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	currSlot := len(nums1) - 1
	ind1 := m - 1
	for ind2 := n - 1; ind2 > -1; {
		if ind1 < 0 || nums2[ind2] >= nums1[ind1] {
			nums1[currSlot] = nums2[ind2]
			ind2--
		} else if nums2[ind2] < nums1[ind1] {
			nums1[currSlot] = nums1[ind1]
			nums1[ind1] = 0
			ind1--
		}
		currSlot--
	}
	fmt.Println(nums1)
}
