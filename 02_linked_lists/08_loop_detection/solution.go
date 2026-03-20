package loopdetection

import "github.com/NoamBechhofer/CTCI-Go/lib"

func LoopDetection[T any](list lib.SinglyLinkedList[T]) *lib.SinglyLinkedListNode[T] {
	if list.Head == nil {
		return nil
	}

	slow := list.Head
	fast := list.Head

	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	ret := list.Head
	for ret != fast {
		ret = ret.Next
		fast = fast.Next
	}
	return ret
}
