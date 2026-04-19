package validatebst

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type testCase struct {
	tree *lib.BinaryTreeNode[int]
	want bool
}

func TestValidateBST(t *testing.T) {
	testCases := []testCase{
		{tree: nil, want: true},
		{tree: &lib.BinaryTreeNode[int]{Val: 1}, want: true},
		{tree: &lib.BinaryTreeNode[int]{Val: 2, Left: &lib.BinaryTreeNode[int]{Val: 1}}, want: true},
		{tree: &lib.BinaryTreeNode[int]{Val: 2, Right: &lib.BinaryTreeNode[int]{Val: 3}}, want: true},
		{tree: &lib.BinaryTreeNode[int]{Val: 2, Left: &lib.BinaryTreeNode[int]{Val: 1}, Right: &lib.BinaryTreeNode[int]{Val: 3}}, want: true},
		{tree: &lib.BinaryTreeNode[int]{Val: 8, Left: &lib.BinaryTreeNode[int]{Val: 4, Left: &lib.BinaryTreeNode[int]{Val: 2}, Right: &lib.BinaryTreeNode[int]{Val: 6}}, Right: &lib.BinaryTreeNode[int]{Val: 12, Left: &lib.BinaryTreeNode[int]{Val: 10}, Right: &lib.BinaryTreeNode[int]{Val: 14}}}, want: true},
		{tree: &lib.BinaryTreeNode[int]{Val: 8, Left: &lib.BinaryTreeNode[int]{Val: 9}}, want: false},
		{tree: &lib.BinaryTreeNode[int]{Val: 8, Right: &lib.BinaryTreeNode[int]{Val: 7}}, want: false},
		{tree: &lib.BinaryTreeNode[int]{Val: 8, Left: &lib.BinaryTreeNode[int]{Val: 4, Right: &lib.BinaryTreeNode[int]{Val: 9}}, Right: &lib.BinaryTreeNode[int]{Val: 12}}, want: false},
		{tree: &lib.BinaryTreeNode[int]{Val: 8, Left: &lib.BinaryTreeNode[int]{Val: 4}, Right: &lib.BinaryTreeNode[int]{Val: 12, Left: &lib.BinaryTreeNode[int]{Val: 6}}}, want: false},
	}

	for i, tc := range testCases {
		t.Run(lib.SignedToString(i), func(t *testing.T) {
			got := ValidateBST(tc.tree)
			if got != tc.want {
				t.Fatalf("want %t, got %t", tc.want, got)
			} else {
				t.Logf("got %t", got)
			}
		})
	}
}
