package removedups

import (
	"fmt"
	"slices"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	list []int32
	want []int32
}

func TestRemoveDups(t *testing.T) {
	testCases := []TestCase{
		{[]int32{}, []int32{}},
		{[]int32{1}, []int32{1}},
		{[]int32{1, 1}, []int32{1}},
		{[]int32{1, 2}, []int32{1, 2}},
		{[]int32{1, 1, 1}, []int32{1}},
		{[]int32{1, 1, 2}, []int32{1, 2}},
		{[]int32{1, 2, 1}, []int32{1, 2}},
		{[]int32{1, 2, 2}, []int32{1, 2}},
		{[]int32{1, 2, 3}, []int32{1, 2, 3}},
	}

	solutions := []struct {
		f    func(*lib.SinglyLinkedList[int32])
		name string
	}{
		{name: "RemoveDups", f: RemoveDups[int32]},
		{name: "RemoveDupsNoTempBuf", f: RemoveDupsNoTempBuf[int32]},
	}

	for _, testCase := range testCases {
		for _, solution := range solutions {
			testName := fmt.Sprintf("%s(%v)", solution.name, testCase.list)
			testFunc := func(t *testing.T) {
				list := lib.SinglyLinkedListFromSlice(testCase.list)
				solution.f(&list)
				got := list.ToSlice()

				if !slices.Equal(got, testCase.want) {
					t.Fatalf("got %v, wanted %v", got, testCase.want)
				} else {
					t.Logf("got %v", got)
				}
			}
			t.Run(testName, testFunc)
		}
	}
}
