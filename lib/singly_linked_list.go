package lib

type SinglyLinkedListNode[T any] struct {
	Val  T
	Next *SinglyLinkedListNode[T]
}

type SinglyLinkedList[T any] struct {
	Head *SinglyLinkedListNode[T]
}

func SinglyLinkedListFromSlice[T any](slice []T) SinglyLinkedList[T] {
	if len(slice) == 0 {
		return SinglyLinkedList[T]{Head: nil}
	}

	block := make([]SinglyLinkedListNode[T], len(slice))

	for i := range len(block) - 1 {
		block[i] = SinglyLinkedListNode[T]{Val: slice[i], Next: &block[i+1]}
	}
	block[len(block)-1] = SinglyLinkedListNode[T]{Val: slice[len(slice)-1], Next: nil}

	return SinglyLinkedList[T]{Head: &block[0]}
}

func (list *SinglyLinkedList[T]) ToSlice() []T {
	ret := []T{}

	for curr := list.Head; curr != nil; curr = curr.Next {
		ret = append(ret, curr.Val)
	}

	return ret
}
