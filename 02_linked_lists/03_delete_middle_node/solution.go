package deletemiddlenode

import "github.com/NoamBechhofer/CTCI-Go/lib"

func DeleteMiddleNode[T any](node *lib.SinglyLinkedListNode[T]) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
	// clean up
	next.Next = nil
}
