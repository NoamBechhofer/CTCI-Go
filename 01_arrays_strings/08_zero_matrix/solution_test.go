package zeromatrix

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

func TestZeroMatrix(t *testing.T) {
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
			[][]int32{{0}},
			[][]int32{{0}},
		},
		{
			[][]int32{
				{1, 2},
				{3, 4},
			},
			[][]int32{
				{1, 2},
				{3, 4},
			},
		},
		{
			[][]int32{
				{0, 2},
				{3, 0},
			},
			[][]int32{
				{0, 0},
				{0, 0},
			},
		},
		{
			[][]int32{
				{1, 0},
				{3, 4},
			},
			[][]int32{
				{0, 0},
				{3, 0},
			},
		},
		{
			[][]int32{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			[][]int32{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			// Single row, no zeros
			[][]int32{{1, 2, 3, 4}},
			[][]int32{{1, 2, 3, 4}},
		},
		{
			// Single row, one zero
			[][]int32{{1, 0, 3, 4}},
			[][]int32{{0, 0, 0, 0}},
		},
		{
			// Single column, no zeros
			[][]int32{
				{1},
				{2},
				{3},
			},
			[][]int32{
				{1},
				{2},
				{3},
			},
		},
		{
			// Single column, one zero
			[][]int32{
				{1},
				{0},
				{3},
			},
			[][]int32{
				{0},
				{0},
				{0},
			},
		},
		{
			// Zero in first row
			[][]int32{
				{1, 2, 0},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int32{
				{0, 0, 0},
				{4, 5, 0},
				{7, 8, 0},
			},
		},
		{
			// Zero in first column
			[][]int32{
				{1, 2, 3},
				{0, 5, 6},
				{7, 8, 9},
			},
			[][]int32{
				{0, 2, 3},
				{0, 0, 0},
				{0, 8, 9},
			},
		},
		{
			// Multiple zeros same row
			[][]int32{
				{1, 0, 3, 0},
				{5, 6, 7, 8},
			},
			[][]int32{
				{0, 0, 0, 0},
				{5, 0, 7, 0},
			},
		},
		{
			// Multiple zeros same column
			[][]int32{
				{1, 2, 3},
				{4, 0, 6},
				{7, 0, 9},
			},
			[][]int32{
				{1, 0, 3},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			// Zeros on diagonal
			[][]int32{
				{0, 2, 3},
				{4, 0, 6},
				{7, 8, 0},
			},
			[][]int32{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			// All zeros
			[][]int32{
				{0, 0},
				{0, 0},
			},
			[][]int32{
				{0, 0},
				{0, 0},
			},
		},
		{
			// Larger matrix, multiple independent zeros
			[][]int32{
				{1, 2, 3, 4},
				{5, 6, 0, 8},
				{9, 10, 11, 12},
				{0, 14, 15, 16},
			},
			[][]int32{
				{0, 2, 0, 4},
				{0, 0, 0, 0},
				{0, 10, 0, 12},
				{0, 0, 0, 0},
			},
		},
		{
			// Rectangular (more rows)
			[][]int32{
				{1, 2},
				{3, 0},
				{5, 6},
				{7, 8},
			},
			[][]int32{
				{1, 0},
				{0, 0},
				{5, 0},
				{7, 0},
			},
		},
		{
			// Rectangular (more columns)
			[][]int32{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 0},
			},
			[][]int32{
				{1, 2, 3, 4, 0},
				{0, 0, 0, 0, 0},
			},
		},
		{
			// Negative values with zero
			[][]int32{
				{-1, -2, 0},
				{-4, -5, -6},
			},
			[][]int32{
				{0, 0, 0},
				{-4, -5, 0},
			},
		},
		{
			// Zero in last cell
			[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			[][]int32{
				{1, 2, 0},
				{4, 5, 0},
				{0, 0, 0},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.matrix), func(t *testing.T) {
			ZeroMatrix(tc.matrix)
			if !matricesEqual(tc.matrix, tc.want) {
				t.Fatalf("wanted %v, got %v", tc.want, tc.matrix)
			} else {
				t.Logf("got %v", tc.matrix)
			}
		})
	}
}
