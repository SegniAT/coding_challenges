package main

func main() {
}

func majorityElementHashMap(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	majorityEl := nums[0]

	for key, value := range count {
		if count[majorityEl] < value {
			majorityEl = key
		}
	}

	return majorityEl
}

func majorityElementMooreVotingAlgorithm(nums []int) int {
	count, candidate := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
