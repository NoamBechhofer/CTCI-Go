package intersection

import "github.com/NoamBechhofer/CTCI-Go/lib"

func Intersection[T any](a, b lib.SinglyLinkedList[T]) *lib.SinglyLinkedListNode[T] {
	if a.Head == nil || b.Head == nil {
		return nil
	}

	aCurr := a.Head
	bCurr := b.Head

	for aCurr != bCurr {
		aCurr = aCurr.Next
		bCurr = bCurr.Next

		if aCurr == nil && bCurr == nil {
			return nil
		}

		if aCurr == nil {
			aCurr = b.Head
		}
		if bCurr == nil {
			bCurr = a.Head
		}
	}

	return aCurr
}
