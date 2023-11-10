package main

import "math"

func main() {

}

// Naive approach
func containsNearbyDuplicateNaive(nums []int, k int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if i != j && nums[i] == nums[j] && int(math.Abs(float64(i-j))) <= k {
				return true
			}
		}
	}

	return false
}

// Hashmap
func containsNearbyDuplicate(nums []int, k int) bool {
	// length := len(nums)
	trackNumsIndex := make(map[int]int)

	for ind, num := range nums {
		if mapInd, ok := trackNumsIndex[num]; ok && ind-mapInd < k {
			return true
		}
		trackNumsIndex[num] = ind
	}
	return false
}
