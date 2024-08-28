package main

func main() {
}

type Value struct {
	ind  int
	temp int
}

func dailyTemperatures(temperatures []int) []int {
	N := len(temperatures)
	ans := make([]int, N)
	stack := []Value{}

	for i, temp := range temperatures {
		// maintain a decreasing monotonic stack
		for len(stack) > 0 && temp > stack[len(stack)-1].temp {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top.ind] = i - top.ind
		}

		stack = append(stack, Value{ind: i, temp: temp})
	}

	return ans
}
