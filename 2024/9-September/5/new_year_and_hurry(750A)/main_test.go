package main

import "testing"

func TestMaxMiceInHole(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		expect int
	}{
		{
			n:      3,
			k:      222,
			expect: 2,
		},
		{
			n:      4,
			k:      190,
			expect: 4,
		},
		{
			n:      7,
			k:      1,
			expect: 7,
		},
	}

	for _, test := range tests {
		res := solve(test.n, test.k)
		if res != test.expect {
			t.Errorf("solve(%d, %d) ; want %d, got %d", test.n, test.k, test.expect, res)
		}
	}
}
