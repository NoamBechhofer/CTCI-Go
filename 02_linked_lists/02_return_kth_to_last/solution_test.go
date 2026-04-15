package returnkthtolast

import (
	"fmt"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	list []int32
	k    int
	want int32
}

func TestReturnKthToLast(t *testing.T) {
	// assumption: last element is 1st-to-last, not 0th-to-last
	testCases := []TestCase{
		{list: []int32{1}, k: 1, want: 1},
		{list: []int32{1, 2}, k: 1, want: 2},
		{list: []int32{1, 2}, k: 2, want: 1},
		{list: []int32{1, 2, 3}, k: 1, want: 3},
		{list: []int32{1, 2, 3}, k: 2, want: 2},
		{list: []int32{1, 2, 3}, k: 3, want: 1},
		{list: []int32{1, 2, 3, 4}, k: 1, want: 4},
		{list: []int32{1, 2, 3, 4}, k: 2, want: 3},
		{list: []int32{1, 2, 3, 4}, k: 3, want: 2},
		{list: []int32{1, 2, 3, 4}, k: 4, want: 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.list), func(t *testing.T) {
			list := lib.ListFromSlice(tc.list)
			got := ReturnKthToLast(list, tc.k)
			if got != tc.want {
				t.Fatalf("wanted %d, got %d", tc.want, got)
			} else {
				t.Logf("got %d", got)
			}
		})
	}
}
