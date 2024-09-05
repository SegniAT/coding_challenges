package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()

	line := strings.Split(scanner.Text(), " ")
	nums := []int{}
	for _, str := range line {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	fmt.Println(maxValidLength(nums))
}

func maxValidLength(nums []int) int {
	N := len(nums)
	left, right := 2, N+2
	maxLength := 0

	if right%2 != 0 {
		right--
	}

	for left <= right {
		mid := left + (right-left)/2
		if mid%2 != 0 {
			mid--
		}

		if validContinuousSubsegment(nums, mid) {
			maxLength = mid
			left = mid + 2
		} else {
			right = mid - 2
		}
	}

	return maxLength
}

func validContinuousSubsegment(nums []int, length int) bool {
	if length%2 != 0 {
		return false
	}

	N := len(nums)
	for start := 0; start+length <= N; start++ {
		mid := start + length/2

		if isUniform(nums[start:mid], 1) && isUniform(nums[mid:start+length], 2) {
			return true
		}

		if isUniform(nums[start:mid], 2) && isUniform(nums[mid:start+length], 1) {
			return true
		}
	}

	return false
}

func isUniform(slice []int, value int) bool {
	for _, num := range slice {
		if num != value {
			return false
		}
	}

	return true
}
