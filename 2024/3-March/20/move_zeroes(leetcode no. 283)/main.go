package main

import "fmt"

func main() {
	inp := []int{1, 0, 4, 2, 0}
	moveZeroes(inp)
	fmt.Println(inp)

	inp = []int{0, 0, 0}
	moveZeroes(inp)
	fmt.Println(inp)

	inp = []int{1, 2, 4}
	moveZeroes(inp)
	fmt.Println(inp)

	inp = []int{1, 2, 0}
	moveZeroes(inp)
	fmt.Println(inp)

	inp = []int{0, 0, 0, 1, 2, 0}
	moveZeroes(inp)
	fmt.Println(inp)
}

func moveZeroes(nums []int) {
	numsLen := len(nums)
	numPtr, zeroPtr := 0, 0

	for {
		for numPtr < numsLen && nums[numPtr] == 0 {
			numPtr++
		}

		for zeroPtr < numsLen && nums[zeroPtr] != 0 {
			zeroPtr++
		}

		if numPtr >= numsLen || zeroPtr >= numsLen {
			break
		}

		if numPtr > zeroPtr {
			nums[numPtr], nums[zeroPtr] = nums[zeroPtr], nums[numPtr]
			numPtr++
			zeroPtr++
		} else {
			numPtr++
		}
	}

}
