package listofdepths

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func isPermutation[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	aElements := map[T]uint{}
	for _, element := range a {
		_, exists := aElements[element]
		if !exists {
			aElements[element] = 1
		} else {
			aElements[element]++
		}
	}

	for _, element := range b {
		freq, exists := aElements[element]
		if !exists || freq == 0 {
			return false
		}
		aElements[element]--
	}

	return true
}

type testCase struct {
	tree *lib.BinaryTreeNode[int]
	want [][]int
}

func TestListOfDepths(t *testing.T) {
	testCases := []testCase{
		{tree: nil, want: [][]int{}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1}, want: [][]int{{1}}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1, Left: &lib.BinaryTreeNode[int]{Val: 2}}, want: [][]int{{1}, {2}}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1, Right: &lib.BinaryTreeNode[int]{Val: 2}}, want: [][]int{{1}, {2}}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1, Left: &lib.BinaryTreeNode[int]{Val: 2}, Right: &lib.BinaryTreeNode[int]{Val: 3}}, want: [][]int{{1}, {2, 3}}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1, Left: &lib.BinaryTreeNode[int]{Val: 2, Left: &lib.BinaryTreeNode[int]{Val: 4}}, Right: &lib.BinaryTreeNode[int]{Val: 3}}, want: [][]int{{1}, {2, 3}, {4}}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1, Left: &lib.BinaryTreeNode[int]{Val: 2, Left: &lib.BinaryTreeNode[int]{Val: 4}}, Right: &lib.BinaryTreeNode[int]{Val: 3, Right: &lib.BinaryTreeNode[int]{Val: 5}}}, want: [][]int{{1}, {2, 3}, {4, 5}}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1, Left: &lib.BinaryTreeNode[int]{Val: 2, Left: &lib.BinaryTreeNode[int]{Val: 4}}, Right: &lib.BinaryTreeNode[int]{Val: 3, Left: &lib.BinaryTreeNode[int]{Val: 5}, Right: &lib.BinaryTreeNode[int]{Val: 6}}}, want: [][]int{{1}, {2, 3}, {4, 5, 6}}},
		{tree: &lib.BinaryTreeNode[int]{Val: 1, Left: &lib.BinaryTreeNode[int]{Val: 2, Left: &lib.BinaryTreeNode[int]{Val: 4}, Right: &lib.BinaryTreeNode[int]{Val: 5}}, Right: &lib.BinaryTreeNode[int]{Val: 3, Left: &lib.BinaryTreeNode[int]{Val: 6}, Right: &lib.BinaryTreeNode[int]{Val: 7}}}, want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
	}

	for i, tc := range testCases {
		t.Run(lib.SignedToString(i), func(t *testing.T) {
			listOfDepths := ListOfDepths(tc.tree)
			got := [][]int{}
			for _, depth := range listOfDepths.ToSlice() {
				got = append(got, depth.ToSlice())
			}

			if len(got) != len(tc.want) {
				t.Fatalf("wanted %d depths, got %d depths", len(tc.want), len(got))
			} else {
				t.Logf("%d depths\n", len(got))
			}

			for i := 0; i < len(got); i++ {
				wantDepth := tc.want[i]
				gotDepth := got[i]

				if !isPermutation(wantDepth, gotDepth) {
					t.Fatalf("wanted %v, got %v", wantDepth, gotDepth)
				} else {
					t.Logf("%d: %v", i, gotDepth)
				}
			}
		})
	}
}
