package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

func summaryRangesUgly(nums []int) []string {
	ln := len(nums)
	var out []string

	start := 0
	for ind := range nums {
		num := nums[ind]
		nextInd := ind + 1

		if nextInd < ln {
			if num+1 != nums[nextInd] {
				if start == ind {
					out = append(out, fmt.Sprintf("%d", nums[start]))
				} else {
					out = append(out, fmt.Sprintf("%d->%d", nums[start], nums[ind]))
				}
				start = nextInd
			} else {
				continue
			}
		} else {
			if start == ind {
				out = append(out, fmt.Sprintf("%d", nums[start]))
			} else {
				out = append(out, fmt.Sprintf("%d->%d", nums[start], nums[ind]))
			}
		}

	}

	return out
}

func summaryRanges(nums []int) []string {
	ln := len(nums)
	var out []string

	if ln == 0 {
		return out
	}

	if ln == 1 {
		out = append(out, fmt.Sprintf("%d", nums[0]))
		return out
	}

	startInd := 0
	for ind := 1; ind < ln; ind++ {
		if nums[ind-1]+1 != nums[ind] {
			endInd := ind - 1
			if startInd == endInd {
				out = append(out, fmt.Sprintf("%d", nums[startInd]))
			} else {
				out = append(out, fmt.Sprintf("%d->%d", nums[startInd], nums[endInd]))
			}
			startInd = ind
		}
	}

	endInd := ln - 1

	if startInd == endInd {
		out = append(out, fmt.Sprintf("%d", nums[startInd]))
	} else {
		out = append(out, fmt.Sprintf("%d->%d", nums[startInd], nums[endInd]))
	}

	return out
}
