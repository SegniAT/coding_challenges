package main

func main() {

}

func productExceptSelfNaive(nums []int) []int {
	length := len(nums)
	products := make([]int, length)

	for i := 0; i < length; i++ {
		products[i] = 1
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i == j {
				continue
			}
			products[i] *= nums[j]
		}
	}

	return products
}

func productExceptSelfBetter(nums []int) []int {
	length := len(nums)
	leftToRight := make([]int, length)
	rightToLeft := make([]int, length)
	result := make([]int, length)

	// left to right
	leftToRight[0] = nums[0]
	for i := 1; i < length; i++ {
		leftToRight[i] = leftToRight[i-1] * nums[i]
	}

	// right to left
	rightToLeft[length-1] = nums[length-1]
	for i := length - 2; i > -1; i-- {
		rightToLeft[i] = rightToLeft[i+1] * nums[i]
	}

	for i := 0; i < length; i++ {
		if i == 0 {
			result[i] = rightToLeft[i+1]
			continue
		}

		if i == length-1 {
			result[i] = leftToRight[i-1]
			continue
		}

		result[i] = leftToRight[i-1] * rightToLeft[i+1]
	}

	return result
}
