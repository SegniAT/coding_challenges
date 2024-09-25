package main

import "fmt"

func main() {
	c := []int{2, 3, 6, 7}
	t := 7

	fmt.Println(combinationSum(c, t))
}
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	curr := []int{}
	count(candidates, target, 0, curr, &res)
	return res
}

// using "ind" to avoid repeated combinations
func count(c []int, t int, ind int, curr []int, res *[][]int) {
	if t == 0 {
		// NOTE: remember that a slice is a pointer to an underlying array!!!
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
	}

	if t < 0 {
		return
	}

	for i := ind; i < len(c); i++ {
		v := c[i]

		// take
		t -= v
		curr = append(curr, v)
		count(c, t, i, curr, res)
		// don't take the next iteration
		t += v
		curr = curr[:len(curr)-1]
	}
}
