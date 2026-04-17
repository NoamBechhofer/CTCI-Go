package checkbalanced

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type testCase struct {
	tree lib.BinaryTreeNode[int]
	want bool
}

func TestCheckBalanced(t *testing.T) {
	testCases := []testCase{
		{tree: lib.BinaryTreeNode[int]{}, want: true},
		{tree: lib.BinaryTreeNode[int]{Left: &lib.BinaryTreeNode[int]{}}, want: true},
		{tree: lib.BinaryTreeNode[int]{Left: &lib.BinaryTreeNode[int]{Left: &lib.BinaryTreeNode[int]{}}}, want: false},
		{tree: lib.BinaryTreeNode[int]{Right: &lib.BinaryTreeNode[int]{}}, want: true},
		{tree: lib.BinaryTreeNode[int]{Right: &lib.BinaryTreeNode[int]{Right: &lib.BinaryTreeNode[int]{}}}, want: false},
	}

	for i, tc := range testCases {
		t.Run(lib.SignedToString(i), func(t *testing.T) {
			got := CheckBalanced(&tc.tree)
			if got != tc.want {
				t.Fatalf("want %t, got %t", tc.want, got)
			} else {
				t.Logf("got %t", got)
			}
		})
	}
}
