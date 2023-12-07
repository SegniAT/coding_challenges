package main

import "testing"

func TestModifyRange(t *testing.T) {
	tests := []struct {
		name      string
		srcRanges [][]int
		mappings  [][]int
		output    [][]int
	}{
		// {
		// 	"All in range",
		// 	[][]int{{11, 15}, {30, 35}, {45, 66}},
		// 	// 		 11 - 15 	 30 - 40	 	42 - 70
		// 	[][]int{{0, 11, 5}, {10, 30, 21}, {20, 44, 27}},
		// 	[][]int{{0, 4}, {10, 15}, {21, 42}},
		// },
		// {
		// 	"All out of range",
		// 	[][]int{{11, 15}, {30, 35}, {45, 66}},
		// 	// 		 80 - 90 	 100 - 110	 	1000 - 2000
		// 	[][]int{{0, 80, 11}, {90, 100, 11}, {4000, 1000, 1001}},
		// 	[][]int{{11, 15}, {30, 35}, {45, 66}},
		// },
		{
			"fertilizer to water",
			[][]int{{81, 94}, {57, 69}},
			[][]int{{49, 53, 8}, {0, 11, 42}, {42, 0, 7}, {57, 7, 4}},
			[][]int{{81, 94}, {53, 56}, {61, 69}},
		},
		{
			"light to temperature",
			[][]int{{74, 87}, {46, 49}, {54, 62}},
			[][]int{{45, 77, 23}, {81, 45, 19}, {68, 64, 13}},
			[][]int{{78, 80}, {45, 55}, {82, 85}, {90, 98}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outRanges := modifyRange(tt.srcRanges, tt.mappings)

			if !sameRanges(outRanges, tt.output) {
				t.Errorf("Expected %v, got %v", tt.output, outRanges)
			}
		})
	}
}

func sameRanges(rangesOne, rangesTwo [][]int) bool {
	lenOne, lenTwo := len(rangesOne), len(rangesTwo)
	if lenOne != lenTwo {
		return false
	}

	for _, rnOne := range rangesOne {
		found := false
		for _, rnTwo := range rangesTwo {
			if rnOne[0] == rnTwo[0] && rnOne[1] == rnTwo[1] {
				found = true
				continue
			}
		}

		if !found {
			return false
		}
	}

	return true
}
