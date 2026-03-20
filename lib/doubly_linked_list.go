package lib

type DoublyLinkedListNode[T any] struct {
	Val  T
	Next *DoublyLinkedListNode[T]
	Prev *DoublyLinkedListNode[T]
}

type DoublyLinkedList[T any] struct {
	Head *DoublyLinkedListNode[T]
	Tail *DoublyLinkedListNode[T]
}

func DoublyLinkedListFromSlice[T any](slice []T) DoublyLinkedList[T] {
	if len(slice) == 0 {
		return DoublyLinkedList[T]{Head: nil, Tail: nil}
	}

	if len(slice) == 1 {
		node := DoublyLinkedListNode[T]{Prev: nil, Val: slice[0], Next: nil}
		return DoublyLinkedList[T]{Head: &node, Tail: &node}
	}

	block := make([]DoublyLinkedListNode[T], len(slice))

	block[0] = DoublyLinkedListNode[T]{Prev: nil, Val: slice[0], Next: &block[1]}

	for i := 1; i < len(block)-1; i++ {
		block[i] = DoublyLinkedListNode[T]{Prev: &block[i-1], Val: slice[i], Next: &block[i+1]}
	}

	block[len(block)-1] = DoublyLinkedListNode[T]{Prev: &block[len(block)-2], Val: slice[len(slice)-1], Next: nil}

	return DoublyLinkedList[T]{Head: &block[0], Tail: &block[len(block)-1]}
}

func (list *DoublyLinkedList[T]) ToSlice() []T {
	ret := []T{}

	for curr := list.Head; curr != nil; curr = curr.Next {
		ret = append(ret, curr.Val)
	}

	return ret
}
