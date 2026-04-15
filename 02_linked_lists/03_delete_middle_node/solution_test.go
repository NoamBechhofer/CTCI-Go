package deletemiddlenode

import (
	"fmt"
	"slices"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	list      []int32
	nodeIndex int
	want      []int32
}

func TestDeleteMiddleNode(t *testing.T) {
	testCases := []TestCase{
		{list: []int32{1, 2, 3}, nodeIndex: 1, want: []int32{1, 3}},
		{list: []int32{1, 2, 3, 4}, nodeIndex: 1, want: []int32{1, 3, 4}},
		{list: []int32{1, 2, 3, 4}, nodeIndex: 2, want: []int32{1, 2, 4}},
		{list: []int32{1, 2, 3, 4, 5}, nodeIndex: 1, want: []int32{1, 3, 4, 5}},
		{list: []int32{1, 2, 3, 4, 5}, nodeIndex: 2, want: []int32{1, 2, 4, 5}},
		{list: []int32{1, 2, 3, 4, 5}, nodeIndex: 3, want: []int32{1, 2, 3, 5}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 1, want: []int32{1, 3, 4, 5, 6}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 2, want: []int32{1, 2, 4, 5, 6}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 3, want: []int32{1, 2, 3, 5, 6}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 4, want: []int32{1, 2, 3, 4, 6}},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%v, %d", tc.list, tc.nodeIndex)
		testFunc := func(t *testing.T) {
			list := lib.SinglyLinkedListFromSlice(tc.list)
			curr := list.Head
			for range tc.nodeIndex {
				curr = curr.Next
			}
			DeleteMiddleNode(curr)
			got := list.ToSlice()
			if !slices.Equal(got, tc.want) {
				t.Fatalf("wanted %v, got %v", tc.want, got)
			} else {
				t.Logf("got %v", got)
			}
		}
		t.Run(testName, testFunc)
	}
}
