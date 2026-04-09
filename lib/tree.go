package lib

import "cmp"

type BinaryTreeNode[T any] struct {
	Val         T
	Left, Right *BinaryTreeNode[T]
}

func (bt *BinaryTreeNode[T]) Height() int {
	if bt == nil {
		return 0
	}
	return 1 + max(bt.Left.Height(), bt.Right.Height())
}

func IsSearchTree[T cmp.Ordered](bt *BinaryTreeNode[T]) bool {
	if bt == nil {
		return true
	}

	if (bt.Left != nil && bt.Left.Val > bt.Val) || (bt.Right != nil && bt.Right.Val < bt.Val) {
		return false
	}

	return IsSearchTree(bt.Left) && IsSearchTree(bt.Right)
}

func BinaryTreeEquals[T comparable](a, b *BinaryTreeNode[T]) bool {
	if a == nil || b == nil {
		return a == b
	}

	return a.Val == b.Val && BinaryTreeEquals(a.Left, b.Left) && BinaryTreeEquals(a.Right, b.Right)
}

type TreeNode[T any] struct {
	Val      T
	Children []TreeNode[T]
}

func (tree *TreeNode[T]) Height() int {
	if tree == nil {
		return 0
	}

	maxChildHeight := 0
	for _, child := range tree.Children {
		childHeight := child.Height()
		if childHeight > maxChildHeight {
			maxChildHeight = childHeight
		}
	}

	return 1 + maxChildHeight
}

func TreeEquals[T comparable](a, b *TreeNode[T]) bool {
	if a == nil || b == nil {
		return a == b
	}

	if a.Val != b.Val {
		return false
	}

	if len(a.Children) != len(b.Children) {
		return false
	}

	for i := 0; i < len(a.Children); i++ {
		if !TreeEquals(&a.Children[i], &b.Children[i]) {
			return false
		}
	}

	return true
}
