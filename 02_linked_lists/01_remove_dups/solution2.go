package removedups

import "github.com/NoamBechhofer/CTCI-Go/lib"

func RemoveDupsNoTempBuf[T comparable](list *lib.SinglyLinkedList[T]) {
	if list.Head == nil {
		return
	}
	for curr := list.Head; curr != nil; curr = curr.Next {
		runner := curr
		for runner.Next != nil {
			if curr.Val == runner.Next.Val {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
	}
}
