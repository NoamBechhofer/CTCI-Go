package removedups

import "github.com/NoamBechhofer/CTCI-Go/lib"

func RemoveDupsNoTempBuf[T comparable](list *lib.TypedList[T]) {
	if list.Len() <= 1 {
		return
	}
	for curr := list.Front(); curr.Next() != nil; curr = curr.Next() {
		runner := curr.Next()
		for runner != nil {
			if curr.Value() != runner.Value() {
				runner = runner.Next()
				continue
			}
			toDelete := runner
			runner = runner.Next()
			list.Remove(toDelete)
		}
	}
}
