package removedups

import "github.com/NoamBechhofer/CTCI-Go/lib"

func RemoveDups[T comparable](list *lib.SinglyLinkedList[T]) {
	if list.Head == nil {
		return
	}

	set := map[T]struct{}{list.Head.Val: {}}

	curr := list.Head
	for curr.Next != nil {
		if _, exists := set[curr.Next.Val]; exists {
			curr.Next = curr.Next.Next
		} else {
			set[curr.Next.Val] = struct{}{}
			curr = curr.Next
		}
	}
}
