package successor

import (
	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func Successor[T any](node *lib.BinaryTreeNode[T]) *lib.BinaryTreeNode[T] {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		ret := node.Right
		for ret.Left != nil {
			ret = ret.Left
		}
		return ret
	}

	ret := node
	for ret.Parent != nil && ret.Parent.Right == ret {
		ret = ret.Parent
	}
	return ret.Parent
}
