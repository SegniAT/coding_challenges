package main

import "sort"

func main() {
}

func smallestDistancePair(nums []int, k int) int {
	N := len(nums)

	sort.Ints(nums)

	minDist, maxDist := 0, nums[N-1]-nums[0]

	for maxDist > minDist {
		//NOTE: using this instead of (min+max)/2 to avoid integer overflow b/c of the summation
		midDist := minDist + (maxDist-minDist)/2

		count := countPairs(nums, midDist)

		if count < k {
			minDist = midDist + 1
		} else {
			maxDist = midDist
		}

	}

	return minDist
}

/*
NOTE: 2 pointer method to find number of pairs with difference *less than or equal to* distance.
NOTE: ON A SORTED ARRAY!!!
NOTE: left is the actual first num of the pair...and right is the upper limit
*/
func countPairs(nums []int, distance int) int {
	count, left := 0, 0
	for right := 1; right < len(nums); right++ {
		for nums[right]-nums[left] > distance {
			left++
		}

		count += right - left
	}

	return count
}
