package returnkthtolast

import (
	"fmt"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	list     []int32
	k        int
	expected int32
}

func TestReturnKthToLast(t *testing.T) {
	// assumption: last element is 1st-to-last, not 0th-to-last
	testCases := []TestCase{
		{list: []int32{1}, k: 1, expected: 1},
		{list: []int32{1, 2}, k: 1, expected: 2},
		{list: []int32{1, 2}, k: 2, expected: 1},
		{list: []int32{1, 2, 3}, k: 1, expected: 3},
		{list: []int32{1, 2, 3}, k: 2, expected: 2},
		{list: []int32{1, 2, 3}, k: 3, expected: 1},
		{list: []int32{1, 2, 3, 4}, k: 1, expected: 4},
		{list: []int32{1, 2, 3, 4}, k: 2, expected: 3},
		{list: []int32{1, 2, 3, 4}, k: 3, expected: 2},
		{list: []int32{1, 2, 3, 4}, k: 4, expected: 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ReturnKthToLast(%v)", tc.list), func(t *testing.T) {
			list := lib.ListFromSlice(tc.list)
			got := ReturnKthToLast(list, tc.k)
			if got != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, got)
			} else {
				t.Logf("got %d", got)
			}
		})
	}
}
