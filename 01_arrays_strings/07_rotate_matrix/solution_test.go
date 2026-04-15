package rotate_matrix

import (
	"fmt"
	"slices"
	"testing"
)

type TestCase struct {
	matrix [][]int32
	want   [][]int32
}

func matricesEqual(m1 [][]int32, m2 [][]int32) bool {
	if len(m1) != len(m2) {
		return false
	}

	for i := range m1 {
		if !slices.Equal(m1[i], m2[i]) {
			return false
		}
	}

	return true
}

func TestRotateMatrix(t *testing.T) {

	testCases := []TestCase{
		{
			[][]int32{},
			[][]int32{},
		},
		{
			[][]int32{{}},
			[][]int32{{}},
		},
		{
			[][]int32{{1}},
			[][]int32{{1}},
		},
		{
			[][]int32{{1, 2}, {3, 4}},
			[][]int32{{3, 1}, {4, 2}},
		},
		{
			[][]int32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int32{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int32{
				{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16},
			},
			[][]int32{
				{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4},
			},
		},
		{
			[][]int32{
				{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20}, {21, 22, 23, 24, 25},
			},
			[][]int32{
				{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4}, {25, 20, 15, 10, 5},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.matrix), func(t *testing.T) {
			RotateMatrix(tc.matrix)
			if !matricesEqual(tc.matrix, tc.want) {
				t.Fatalf("wanted %v, got %v", tc.want, tc.matrix)
			} else {
				t.Logf("%v", tc.matrix)
			}
		})
	}
}
