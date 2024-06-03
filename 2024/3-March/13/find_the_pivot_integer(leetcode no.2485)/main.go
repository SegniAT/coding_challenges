package main

import "fmt"

func main() {
	fmt.Println(pivotInteger(8))
	fmt.Println(pivotInteger(1))
	fmt.Println(pivotInteger(4))
}

func pivotInteger(n int) int {
	prefixSum := []int{}

	for i := range n {
		prefixSum = append(prefixSum, i+1)
	}

	for i := 1; i < n; i++ {
		prefixSum[i] += prefixSum[i-1]
	}

	fmt.Println(prefixSum)

	for i := range n {
		before, after := prefixSum[i], prefixSum[n-1]-prefixSum[i]+(i+1)
		//	fmt.Println(before, after)
		if before == after {
			return i + 1
		}
	}

	return -1
}
