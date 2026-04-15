package palindrome

import "github.com/NoamBechhofer/CTCI-Go/lib"

func Palindrome[T comparable](list lib.SinglyLinkedList[T]) bool {
	if list.Head == nil || list.Head.Next == nil {
		return true
	}

	stack := lib.ArrayStack[T]{}

	slow := list.Head
	fast := list.Head

	for fast != nil {
		if fast.Next == nil {
			// odd
			slow = slow.Next
			break
		}
		stack.Push(slow.Val)
		fast = fast.Next.Next
		slow = slow.Next
	}

	for ; slow != nil; slow = slow.Next {
		pop, _ := stack.Pop()
		if slow.Val != pop {
			return false
		}
	}

	return true
}

func getLength[T any](node *lib.SinglyLinkedListNode[T]) int {
	if node == nil {
		return 0
	}

	return 1 + getLength(node.Next)
}

func PalindromeRecursive[T comparable](list lib.SinglyLinkedList[T]) bool {
	listLen := getLength(list.Head)
	_, ret := palindromeRecursiveDriver(list.Head, listLen)
	return ret
}

func palindromeRecursiveDriver[T comparable](node *lib.SinglyLinkedListNode[T], listLen int) (*lib.SinglyLinkedListNode[T], bool) {
	if listLen == 0 {
		return node, true
	}
	if listLen == 1 {
		return node.Next, true
	}

	correspondingNode, recursionIsPalindrome := palindromeRecursiveDriver(node.Next, listLen-2)

	isPalindrome := recursionIsPalindrome && node.Val == correspondingNode.Val

	return correspondingNode.Next, isPalindrome
}
