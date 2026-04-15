package loopdetection

import (
	"fmt"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	list      []int32
	loopIndex int // -1 indicates no loop
}

func (tc *TestCase) generateList() (lib.SinglyLinkedList[int32], *lib.SinglyLinkedListNode[int32]) {
	if tc.loopIndex < 0 {
		return lib.SinglyLinkedListFromSlice(tc.list), nil
	}
	if len(tc.list) == 0 {
		return lib.SinglyLinkedList[int32]{Head: nil}, nil
	}

	var loopNode, next, tail *lib.SinglyLinkedListNode[int32]

	next = nil
	for i := len(tc.list) - 1; i >= 0; i-- {
		next = &lib.SinglyLinkedListNode[int32]{Val: tc.list[i], Next: next}
		if i == tc.loopIndex {
			loopNode = next
		}
		if i == len(tc.list)-1 {
			tail = next
		}
	}
	tail.Next = loopNode

	return lib.SinglyLinkedList[int32]{Head: next}, loopNode
}

func TestLoopDetection(t *testing.T) {
	testCases := []TestCase{
		{list: []int32{}, loopIndex: -1},
		{list: []int32{1}, loopIndex: -1},
		{list: []int32{1}, loopIndex: 0},
		{list: []int32{1, 2}, loopIndex: -1},
		{list: []int32{1, 2}, loopIndex: 0},
		{list: []int32{1, 2}, loopIndex: 1},
		{list: []int32{1, 2, 3}, loopIndex: -1},
		{list: []int32{1, 2, 3}, loopIndex: 0},
		{list: []int32{1, 2, 3}, loopIndex: 1},
		{list: []int32{1, 2, 3}, loopIndex: 2},
		{list: []int32{1, 2, 3, 4}, loopIndex: -1},
		{list: []int32{1, 2, 3, 4}, loopIndex: 0},
		{list: []int32{1, 2, 3, 4}, loopIndex: 1},
		{list: []int32{1, 2, 3, 4}, loopIndex: 2},
		{list: []int32{1, 2, 3, 4}, loopIndex: 3},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%v, %d", tc.list, tc.loopIndex)
		testFunc := func(t *testing.T) {
			list, want := tc.generateList()

			got := LoopDetection(list)
			var gotVal string
			if got != nil {
				gotVal = lib.SignedToString(got.Val)
			} else {
				gotVal = "nil"
			}

			if want != got {
				var expectedVal string
				if want != nil {
					expectedVal = lib.SignedToString(want.Val)
				} else {
					expectedVal = "nil"
				}

				t.Fatalf("wanted %p (%s), got %p (%s)", want, expectedVal, got, gotVal)
			} else {
				t.Logf("got %p (%s)", got, gotVal)
			}
		}
		t.Run(testName, testFunc)
	}
}
