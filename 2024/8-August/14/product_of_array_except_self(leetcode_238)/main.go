package main

func main() {
}

func productExceptSelf(nums []int) []int {
	N := len(nums)
	prefixProductLtR, prefixProductRtL := make([]int, N), make([]int, N)
	copy(prefixProductLtR, nums)
	copy(prefixProductRtL, nums)

	for i := range nums {
		if i-1 < 0 {
			continue
		}
		prefixProductLtR[i] = prefixProductLtR[i] * prefixProductLtR[i-1]
	}

	for i := N - 2; i > -1; i-- {
		prefixProductRtL[i] = prefixProductRtL[i] * prefixProductRtL[i+1]
	}

	res := make([]int, N)

	for i := range nums {
		left := 1
		if i-1 > -1 {
			left = prefixProductLtR[i-1]
		}

		right := 1
		if i+1 < N {
			right = prefixProductRtL[i+1]
		}

		res[i] = left * right
	}

	return res
}
