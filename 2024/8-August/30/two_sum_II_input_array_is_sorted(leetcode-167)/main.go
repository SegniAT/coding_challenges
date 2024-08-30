package main

func main() {
}

// always has a solution according to the constraints
func twoSum(numbers []int, target int) []int {
	N := len(numbers)
	left, right := 0, N-1

	// always has a solution according to the constraints
	for sum := numbers[left] + numbers[right]; sum != target; sum = numbers[left] + numbers[right] {
		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}
