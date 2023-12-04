package main

func main() {

}

func rotate(nums []int, k int) {
	numsLen := len(nums)
	res := make([]int, numsLen)

	for ind, num := range nums {
		nextInd := (ind + k) % numsLen
		res[nextInd] = num
	}
	nums = res
}

func rotate2(nums []int, k int) {
	length := len(nums)
	currNum := nums[0]
	currInd := 0

	for operations := 0; operations < len(nums); operations++ {
		nextIndex := (currInd + k) % length
		temp := nums[nextIndex]
		nums[nextIndex] = currNum

		currNum = temp
		currInd = nextIndex
	}
}
