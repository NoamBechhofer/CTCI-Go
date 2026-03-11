package partition

import (
	"cmp"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func Partition[T cmp.Ordered](list *lib.SinglyLinkedList[T], partitionValue T) {
	if list.Head == nil {
		return
	}

	// first node not yet confirmed to be in the left (AKA < partitionValue) partition
	part := list.Head
	runner := list.Head

	for runner != nil {
		for part != nil && part.Val < partitionValue {
			if runner == part {
				runner = runner.Next
			}
			part = part.Next
		}
		if part == nil {
			return
		}
		for runner != nil && runner.Val >= partitionValue {
			runner = runner.Next
		}
		if runner == nil {
			return
		}
		part.Val, runner.Val = runner.Val, part.Val
		part = part.Next
		runner = runner.Next
	}
}
