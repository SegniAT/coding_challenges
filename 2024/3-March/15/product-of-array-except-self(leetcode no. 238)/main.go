package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println()
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	ln := len(nums)
	leftToRight, rightToLeft := make([]int, ln), make([]int, ln)
	ans := make([]int, ln)

	copy(leftToRight, nums)
	copy(rightToLeft, nums)

	for i := 1; i < len(nums); i++ {
		leftToRight[i] *= leftToRight[i-1]
	}

	for i := len(nums) - 2; i > 0; i-- {
		rightToLeft[i] *= rightToLeft[i+1]
	}

	fmt.Println("ltr", leftToRight)
	fmt.Println("rtl", rightToLeft)

	for i := 0; i < len(nums); i++ {
		left, right := 1, 1
		if i > 0 {
			left = leftToRight[i-1]
		}

		if i < len(nums)-1 {
			right = rightToLeft[i+1]
		}

		ans[i] = left * right
	}
	return ans
}
