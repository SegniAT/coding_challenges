package main

import (
	"fmt"
)

func main() {
	nums, minK, maxK := []int{1, 3, 5, 2, 7, 5}, 1, 5
	fmt.Println(countSubarrays(nums, minK, maxK))

	nums, minK, maxK = []int{1, 1, 1, 1}, 1, 1
	fmt.Println(countSubarrays(nums, minK, maxK))

	nums, minK, maxK = []int{978_650, 978_650, 978_650, 68_071, 52_201, 68_071, 186_141, 978_650, 978_650, 267_135, 68_071, 717_241, 242_895, 68_071, 582_505, 978_650, 68_071, 68_071}, 68_071, 978_650
	fmt.Println(countSubarrays(nums, minK, maxK))
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	localMinPtr, localMaxPtr := -1, -1
	leftPtr := 0
	var totalSubarrays int64

	for rightPtr, num := range nums {
		if num < minK || num > maxK {
			localMinPtr, localMaxPtr = -1, -1
			leftPtr = rightPtr + 1
		} else {
			if num == minK {
				localMinPtr = rightPtr
			}

			if num == maxK {
				localMaxPtr = rightPtr
			}

			totalSubarrays += int64(max(0, min(localMaxPtr, localMinPtr)-leftPtr+1))

		}
	}

	return totalSubarrays

}
