package minimaltree

import (
	"github.com/NoamBechhofer/CTCI-Go/lib"
	"golang.org/x/exp/constraints"
)

func MinimalTree[T constraints.Integer](uniqueIncreasing []T) *lib.BinaryTreeNode[T] {
	if len(uniqueIncreasing) == 0 {
		return nil
	}

	midIdx := len(uniqueIncreasing) / 2
	return &lib.BinaryTreeNode[T]{
		Val:   uniqueIncreasing[midIdx],
		Left:  MinimalTree(uniqueIncreasing[:midIdx]),
		Right: MinimalTree(uniqueIncreasing[midIdx+1:]),
	}
}
