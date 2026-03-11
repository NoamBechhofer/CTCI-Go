package partition

import (
	"fmt"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	list []int32
}

func isValidPartition(list []int32, partitionValue int32) bool {
	if len(list) == 0 {
		return true
	}
findPartition:
	for i := range list {
		for j := 0; j < i; j++ {
			if list[j] >= partitionValue {
				continue findPartition
			}
		}
		for j := i; j < len(list); j++ {
			if list[j] < partitionValue {
				continue findPartition
			}
		}
		return true
	}
	return false
}

func TestPartition(t *testing.T) {
	testCases := []TestCase{
		{list: []int32{}},
		{list: []int32{0}},
		{list: []int32{0, 1}},
		{list: []int32{1, 0}},
		{list: []int32{3, 5, 8, 5, 10, 2, 1}},
	}

	for _, tc := range testCases {
		for _, partitionValue := range tc.list {
			testName := fmt.Sprintf("Partition(%v, %d)", tc.list, partitionValue)
			testFunc := func(t *testing.T) {
				list := lib.SinglyLinkedListFromSlice(tc.list)
				Partition(&list, partitionValue)
				got := list.ToSlice()
				t.Logf("got %v", got)
				if !isValidPartition(got, partitionValue) {
					t.FailNow()
				}
			}
			t.Run(testName, testFunc)
		}

	}
}
