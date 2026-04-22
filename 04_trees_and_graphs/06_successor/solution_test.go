package successor

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func assert(t *testing.T, want, got *lib.BinaryTreeNode[int]) {
	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	} else {
		t.Logf("got %v", got)
	}
}

func TestSuccessor(t *testing.T) {
	t.Run("nil tree", func(t *testing.T) {
		got := Successor[int](nil)
		want := (*lib.BinaryTreeNode[int])(nil)
		assert(t, want, got)
	})

	root := lib.BinaryTreeNode[int]{Val: 0}
	/*
		0
	*/
	t.Run("singleton tree", func(t *testing.T) {
		got := Successor(&root)
		want := (*lib.BinaryTreeNode[int])(nil)
		assert(t, want, got)
	})

	root.Left = &lib.BinaryTreeNode[int]{Val: -1_000, Parent: &root}
	/*
				0
			   /
		-1_000
	*/
	t.Run("root and one left child", func(t *testing.T) {
		got := Successor(root.Left)
		want := &root
		assert(t, want, got)
	})

	root.Left.Left = &lib.BinaryTreeNode[int]{Val: -1_500, Parent: root.Left}
	/*
					  0
					 /
			   -1_000
			  /
		-1_500
	*/
	t.Run("degenerate leftward linked list", func(t *testing.T) {
		got := Successor(root.Left.Left)
		want := root.Left
		assert(t, want, got)
	})

	root.Left.Right = &lib.BinaryTreeNode[int]{Val: -500, Parent: root.Left}
	/*
					   0
					 /
			   -1_000
			  /      \
		-1_500        -500
	*/
	t.Run("root's left child has two children", func(t *testing.T) {
		got := Successor(root.Left)
		want := root.Left.Right
		assert(t, want, got)
	})

	root.Left.Right.Left = &lib.BinaryTreeNode[int]{Val: -750, Parent: root.Left.Right}
	/*
					   0
					 /
			   -1_000
			  /      \
		-1_500        -500
					 /
				-750
	*/
	t.Run("root's left child's right child has a left child", func(t *testing.T) {
		got := Successor(root.Left)
		want := root.Left.Right.Left
		assert(t, want, got)
	})

	root.Left.Right.Right = &lib.BinaryTreeNode[int]{Val: -250, Parent: root.Left.Right}
	/*
					   0
					 /
			   -1_000
			  /      \
		-1_500        -500
					 /	  \
				-750		-250
	*/
	t.Run("root's left child's right child has two children", func(t *testing.T) {
		got := Successor(root.Left.Right.Right)
		want := &root
		assert(t, want, got)
	})

	root.Right = &lib.BinaryTreeNode[int]{Val: 1_000, Parent: &root}
	/*
					   0
					 /	 \
			   -1_000	   1_000
			  /      \
		-1_500        -500
					 /	  \
				-750		-250
	*/
	t.Run("root's right child is a leaf", func(t *testing.T) {
		got := Successor(root.Right)
		want := (*lib.BinaryTreeNode[int])(nil)
		assert(t, want, got)
	})

	root.Right.Right = &lib.BinaryTreeNode[int]{Val: 1_500, Parent: root.Right}
	/*
					   0
					 /	 \
			   -1_000	   1_000
			  /      \			 \
		-1_500        -500		   1_500
					 /	  \
				-750		-250
	*/
	t.Run("root's right child has a right child", func(t *testing.T) {
		got := Successor(root.Right)
		want := root.Right.Right
		assert(t, want, got)
	})

	root.Right.Left = &lib.BinaryTreeNode[int]{Val: 500, Parent: root.Right}
	/*
						 0
					 /		  \
			   -1_000			1_000
			  /      \		   /	 \
		-1_500        -500  500		  1_500
					 /	  \
				-750		-250
	*/
	t.Run("root's right child has two children", func(t *testing.T) {
		got := Successor(root.Right.Left)
		want := root.Right
		assert(t, want, got)
	})

	t.Run("maximum node has no successor", func(t *testing.T) {
		got := Successor(root.Right.Right)
		want := (*lib.BinaryTreeNode[int])(nil)
		assert(t, want, got)
	})

	t.Run("root successor is leftmost of right subtree", func(t *testing.T) {
		got := Successor(&root)
		want := root.Right.Left
		assert(t, want, got)
	})

	t.Run("leftmost node successor", func(t *testing.T) {
		got := Successor(root.Left.Left)
		want := root.Left
		assert(t, want, got)
	})
}
