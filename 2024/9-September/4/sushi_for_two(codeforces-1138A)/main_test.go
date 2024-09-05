package main

import (
	"testing"
)

func TestValidContinuousSubsegment(t *testing.T) {
	tests := []struct {
		nums   []int
		length int
		expect bool
	}{
		{[]int{1, 1, 2, 2}, 4, true},        // Valid segment [1, 1, 2, 2]
		{[]int{2, 2, 1, 1}, 4, true},        // Valid segment [2, 2, 1, 1]
		{[]int{1, 2, 1, 2}, 2, true},        // Valid segments [1, 2] and [2, 1]
		{[]int{1, 2, 2, 1, 1}, 2, true},     // Valid segments [1, 2] and [2, 1]
		{[]int{1, 1, 1, 2, 2, 2}, 6, true},  // Valid segment [1, 1, 1, 2, 2, 2]
		{[]int{1, 1, 2, 2, 1, 1}, 4, true},  // Valid segments [1, 1, 2, 2] and [2, 2, 1, 1]
		{[]int{1, 2, 1, 2, 1, 2}, 4, false}, // No valid segment of length 4
		{[]int{1, 2}, 2, true},              // Valid segment [1, 2]
		{[]int{2, 1}, 2, true},              // Valid segment [2, 1]
		{[]int{1, 1, 2}, 2, true},           // Valid segment [1, 2]
	}

	for _, test := range tests {
		result := validContinuousSubsegment(test.nums, test.length)
		if result != test.expect {
			t.Errorf("validContinuousSubsegment(%v, %d) = %v; want %v", test.nums, test.length, result, test.expect)
		}
	}
}

func TestMaxValidLength(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{1, 2, 1, 2}, 2},       // Valid segment [1, 2,] or [1, 2]
		{[]int{2, 2, 1, 1}, 4},       // Valid segment [2, 2, 1, 1]
		{[]int{1, 1, 2, 2, 2, 1}, 4}, // Valid segment [1, 1, 2, 2]
		{[]int{1, 2, 2, 1}, 2},       // Valid segments [1, 2] and [2, 1]
		{[]int{1, 2}, 2},             // Valid segment [1, 2]
		{[]int{1, 1, 1, 1}, 0},       // No valid segment
		{[]int{1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}, 10},
	}

	for _, test := range tests {
		result := maxValidLength(test.nums)
		if result != test.expect {
			t.Errorf("maxValidLength(%v) = %d; want %d", test.nums, result, test.expect)
		}
	}
}
