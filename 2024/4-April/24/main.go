package main

func main() {
}

func tribonnacci(n int) int {
	memo := make(map[int]int)
	return helper(n, &memo)
}

func helper(n int, memo *map[int]int) int {
	if n == 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	if val, ok := (*memo)[n]; ok {
		return val
	}

	res := helper(n-1, memo) + helper(n-2, memo) + helper(n-3, memo)
	(*memo)[n] = res

	return res
}
