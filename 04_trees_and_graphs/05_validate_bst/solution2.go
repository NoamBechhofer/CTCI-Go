package validatebst

import (
	"cmp"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func ValidateBST[T cmp.Ordered](tree *lib.BinaryTreeNode[T]) bool {
	var zero T
	return validateBstDriver(tree, zero, zero, false, false)
}

// min is exclusive, max is inclusive.
func validateBstDriver[T cmp.Ordered](node *lib.BinaryTreeNode[T], min, max T, haveMin, haveMax bool) bool {
	if node == nil {
		return true
	}

	if (haveMin && node.Val <= min) || (haveMax && node.Val > max) {
		return false
	}

	return validateBstDriver(node.Right, node.Val, max, true, haveMax) && validateBstDriver(node.Left, min, node.Val, haveMin, true)
}
