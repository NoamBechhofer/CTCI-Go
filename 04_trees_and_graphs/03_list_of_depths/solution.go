package listofdepths

import "github.com/NoamBechhofer/CTCI-Go/lib"

func ListOfDepths[T any](bt *lib.BinaryTreeNode[T]) lib.SinglyLinkedList[lib.SinglyLinkedList[T]] {
	if bt == nil {
		return lib.SinglyLinkedList[lib.SinglyLinkedList[T]]{}
	}

	nextLevel := &lib.ArrayQueue[*lib.BinaryTreeNode[T]]{}
	nextLevel.Add(bt)

	ret := lib.SinglyLinkedList[lib.SinglyLinkedList[T]]{}
	retTail := ret.Head

	for !nextLevel.IsEmpty() {
		currentLevel := nextLevel
		nextLevel = &lib.ArrayQueue[*lib.BinaryTreeNode[T]]{}

		list := lib.SinglyLinkedList[T]{}
		listTail := list.Head

		for !currentLevel.IsEmpty() {
			ele, _ := currentLevel.Remove()

			newListTail := lib.SinglyLinkedListNode[T]{Val: ele.Val, Next: nil}
			if listTail == nil {
				list.Head = &newListTail
			} else {
				listTail.Next = &newListTail
			}
			listTail = &newListTail

			if ele.Left != nil {
				nextLevel.Add(ele.Left)
			}
			if ele.Right != nil {
				nextLevel.Add(ele.Right)
			}
		}

		newRetTail := lib.SinglyLinkedListNode[lib.SinglyLinkedList[T]]{Val: list, Next: nil}
		if retTail == nil {
			ret.Head = &newRetTail
		} else {
			retTail.Next = &newRetTail
		}
		retTail = &newRetTail
	}

	return ret
}
