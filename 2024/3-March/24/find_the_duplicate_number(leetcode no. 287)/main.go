package main

func main() {
}

func findDuplicate(nums []int) int {
	hashMap := make(map[int]bool)

	for _, num := range nums {
		if _, ok := hashMap[num]; ok {
			return num
		} else {
			hashMap[num] = true
		}
	}

	return -1
}
