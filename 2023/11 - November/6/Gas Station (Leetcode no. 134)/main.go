package main

import "fmt"

func main() {
	// gas := []int{5, 8, 2, 8}
	// cost := []int{6, 5, 6, 6}
	// gas := []int{1,2,3,4,5}
	// cost := []int{3,4,5,1,2}
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuitNaive(gas, cost))
}

func canCompleteCircuitNaive(gas []int, cost []int) int {
	length := len(gas)

	for i := 0; i < length; i++ {
		currGas := 0
		for j := i; j < length*2; j++ {
			currGas += gas[j%length] - cost[j%length]
			if currGas < 0 {
				break
			}

			if i != j && i == j%length {
				return i
			}
		}
	}

	return -1
}

func canCompleteCircuitGreedy(gas []int, cost []int) int {
	length, currGas, currStartingIdx := len(gas), 0, 0

	for i := currStartingIdx; i < length*2; i++ {
		currGas += gas[i%length] - cost[i%length]
		if currGas < 0 {
			currStartingIdx++
			currGas = 0
			continue
		}

		if currStartingIdx != i && currStartingIdx == i%length {
			return currStartingIdx
		}
	}

	return -1
}
