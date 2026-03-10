package removedups

import (
	"fmt"
	"slices"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	list     []int32
	expected []int32
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
		f    func(*lib.TypedList[int32])
		name string
	}{
		{name: "RemoveDups", f: RemoveDups[int32]},
		{name: "RemoveDupsNoTempBuf", f: RemoveDupsNoTempBuf[int32]},
	}

	for _, testCase := range testCases {
		for _, solution := range solutions {
			testName := fmt.Sprintf("%s(%v)", solution.name, testCase.list)
			testFunc := func(t *testing.T) {
				list := lib.ListFromSlice(testCase.list)
				solution.f(list)
				got := lib.ListToSlice(list)

				if !slices.Equal(got, testCase.expected) {
					t.Fatalf("got %v, expected %v", got, testCase.expected)
				} else {
					t.Logf("got %v", got)
				}
			}
			t.Run(testName, testFunc)
		}
	}
}
