package removedups

import "github.com/NoamBechhofer/CTCI-Go/lib"

func RemoveDups[T comparable](list *lib.TypedList[T]) {
	if list.Len() <= 1 {
		return
	}

	set := make(map[T]struct{})
	curr := list.Front()
	for curr != nil {
		if _, exists := set[curr.Value()]; exists {
			next := curr.Next()
			list.Remove(curr)
			curr = next
		} else {
			set[curr.Value()] = struct{}{}
			curr = curr.Next()
		}
	}
}
