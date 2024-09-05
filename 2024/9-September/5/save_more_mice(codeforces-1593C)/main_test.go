package main

import "testing"

func TestMaxMiceInHole(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		mice   []int
		expect int
	}{
		{
			n:      10,
			k:      6,
			mice:   []int{8, 7, 5, 4, 9, 4},
			expect: 3,
		},
		{
			n:      2,
			k:      8,
			mice:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			expect: 1,
		},
		{
			n:      12,
			k:      11,
			mice:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			expect: 4,
		},
	}

	for _, test := range tests {
		res := maxMiceInHole(test.n, test.k, test.mice)
		if res != test.expect {
			t.Errorf("maxMiceInHole(%d, %d, %v); want %d, got %d", test.n, test.k, test.mice, test.expect, res)
		}
	}
}
