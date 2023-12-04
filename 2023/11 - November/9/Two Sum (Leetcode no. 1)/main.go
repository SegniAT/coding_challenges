package main

func main() {

}

// Naive
func twoSumNaive(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

// HashMap
func twoSumHashMap(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for ind, num := range nums {
		numsMap[num] = ind
	}

	for ind, num := range nums {
		if ind2, ok := numsMap[target-num]; ok && ind != ind2 {
			return []int{ind, ind2}
		}
	}

	return []int{-1, -1}
}
