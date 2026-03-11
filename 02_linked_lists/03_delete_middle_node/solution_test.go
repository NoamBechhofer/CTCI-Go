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
	expected  []int32
}

func TestDeleteMiddleNode(t *testing.T) {
	testCases := []TestCase{
		{list: []int32{1, 2, 3}, nodeIndex: 1, expected: []int32{1, 3}},
		{list: []int32{1, 2, 3, 4}, nodeIndex: 1, expected: []int32{1, 3, 4}},
		{list: []int32{1, 2, 3, 4}, nodeIndex: 2, expected: []int32{1, 2, 4}},
		{list: []int32{1, 2, 3, 4, 5}, nodeIndex: 1, expected: []int32{1, 3, 4, 5}},
		{list: []int32{1, 2, 3, 4, 5}, nodeIndex: 2, expected: []int32{1, 2, 4, 5}},
		{list: []int32{1, 2, 3, 4, 5}, nodeIndex: 3, expected: []int32{1, 2, 3, 5}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 1, expected: []int32{1, 3, 4, 5, 6}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 2, expected: []int32{1, 2, 4, 5, 6}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 3, expected: []int32{1, 2, 3, 5, 6}},
		{list: []int32{1, 2, 3, 4, 5, 6}, nodeIndex: 4, expected: []int32{1, 2, 3, 4, 6}},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("DeleteMiddleNode(%v, %d)", tc.list, tc.nodeIndex)
		testFunc := func(t *testing.T) {
			list := lib.SinglyLinkedListFromSlice(tc.list)
			curr := list.Head
			for range tc.nodeIndex {
				curr = curr.Next
			}
			DeleteMiddleNode(curr)
			got := list.ToSlice()
			if !slices.Equal(got, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, got)
			} else {
				t.Logf("got %v", got)
			}
		}
		t.Run(testName, testFunc)
	}
}
