package returnkthtolast

import "github.com/NoamBechhofer/CTCI-Go/lib"

func ReturnKthToLast[T any](list *lib.TypedList[T], k int) T {
	right := list.Front()
	left := list.Front()

	for i := 1; right.Next() != nil; i++ {

		right = right.Next()
		if i >= k {
			left = left.Next()
		}
	}

	return left.Value()
}
