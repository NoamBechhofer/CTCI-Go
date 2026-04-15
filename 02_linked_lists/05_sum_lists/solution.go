package sumlists

import (
	"github.com/NoamBechhofer/CTCI-Go/lib"
	"golang.org/x/exp/constraints"
)

func advanceList[T constraints.Unsigned](list *lib.SinglyLinkedListNode[T]) *lib.SinglyLinkedListNode[T] {
	if list == nil {
		return nil
	}
	return list.Next
}

func valOrZero[T constraints.Unsigned](list *lib.SinglyLinkedListNode[T]) T {
	if list == nil {
		return 0
	}
	return list.Val
}

func normalizeForwardLists[T constraints.Unsigned](
	a lib.SinglyLinkedList[T], b lib.SinglyLinkedList[T],
) (lib.SinglyLinkedList[T], lib.SinglyLinkedList[T]) {
	aPadding, bPadding := normalizeForwardListsReport(a.Head, b.Head)
	newA := a.Head
	newB := b.Head
	for range aPadding {
		newA = &lib.SinglyLinkedListNode[T]{Val: 0, Next: newA}
	}
	for range bPadding {
		newB = &lib.SinglyLinkedListNode[T]{Val: 0, Next: newB}
	}
	return lib.SinglyLinkedList[T]{Head: newA}, lib.SinglyLinkedList[T]{Head: newB}
}

func normalizeForwardListsReport[T constraints.Unsigned](
	a *lib.SinglyLinkedListNode[T], b *lib.SinglyLinkedListNode[T],
) (int, int) {
	aPadding := 0
	bPadding := 0
	for a != nil || b != nil {
		if a == nil {
			aPadding++
		} else {
			a = a.Next
		}
		if b == nil {
			bPadding++
		} else {
			b = b.Next
		}
	}
	return aPadding, bPadding
}

func SumListsForward[T constraints.Unsigned](a lib.SinglyLinkedList[T], b lib.SinglyLinkedList[T]) lib.SinglyLinkedList[T] {
	newA, newB := normalizeForwardLists(a, b)
	head, carryover := sumListsForwardDriver(newA.Head, newB.Head)
	if carryover != 0 {
		head = &lib.SinglyLinkedListNode[T]{Val: carryover, Next: head}
	}
	return lib.SinglyLinkedList[T]{Head: head}
}

// invariant: len(a) == len(b)
func sumListsForwardDriver[T constraints.Unsigned](a *lib.SinglyLinkedListNode[T], b *lib.SinglyLinkedListNode[T]) (*lib.SinglyLinkedListNode[T], T) {
	if (a == nil) && (b == nil) {
		return nil, 0
	}

	tail, tailCarryover := sumListsForwardDriver(advanceList(a), advanceList(b))

	sum := valOrZero(a) + valOrZero(b) + tailCarryover

	onesPlace := sum % 10
	carryover := sum / 10

	ourNode := lib.SinglyLinkedListNode[T]{Val: onesPlace, Next: tail}

	return &ourNode, carryover
}

func SumListsReverse[T constraints.Unsigned](a lib.SinglyLinkedList[T], b lib.SinglyLinkedList[T]) lib.SinglyLinkedList[T] {
	head := sumListsReverseDriver(a.Head, b.Head, 0)

	return lib.SinglyLinkedList[T]{Head: head}
}

func sumListsReverseDriver[T constraints.Unsigned](a *lib.SinglyLinkedListNode[T], b *lib.SinglyLinkedListNode[T], carryover T) *lib.SinglyLinkedListNode[T] {
	if a == nil && b == nil && carryover == 0 {
		return nil
	}

	sum := carryover
	if a != nil {
		sum += a.Val
	}
	if b != nil {
		sum += b.Val
	}

	onesPlace := sum % 10
	nextCarryover := sum / 10

	tail := sumListsReverseDriver(advanceList(a), advanceList(b), nextCarryover)
	node := lib.SinglyLinkedListNode[T]{Val: onesPlace, Next: tail}
	return &node
}
