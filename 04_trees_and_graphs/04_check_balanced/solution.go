package checkbalanced

import "github.com/NoamBechhofer/CTCI-Go/lib"

func CheckBalanced[T any](tree *lib.BinaryTreeNode[T]) bool {
	_, isBalanced := checkBalancedDriver(tree)
	return isBalanced
}

func checkBalancedDriver[T any](node *lib.BinaryTreeNode[T]) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftOk := checkBalancedDriver(node.Left)
	if !leftOk {
		return 0, false
	}

	rightHeight, rightOk := checkBalancedDriver(node.Right)
	if !rightOk {
		return 0, false
	}

	if abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}

	return 1 + max(leftHeight, rightHeight), true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
