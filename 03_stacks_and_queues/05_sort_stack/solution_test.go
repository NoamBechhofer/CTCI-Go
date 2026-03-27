package sortstack

import (
	"fmt"
	"slices"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	initStack []int
	want      []int
}

func TestSortStack(t *testing.T) {
	testCases := []TestCase{
		{initStack: []int{}, want: []int{}},
		{initStack: []int{1}, want: []int{1}},
		{initStack: []int{1, 2}, want: []int{2, 1}},
		{initStack: []int{2, 1}, want: []int{2, 1}},
		{initStack: []int{1, 2, 3}, want: []int{3, 2, 1}},
		{initStack: []int{1, 3, 2}, want: []int{3, 2, 1}},
		{initStack: []int{2, 1, 3}, want: []int{3, 2, 1}},
		{initStack: []int{2, 3, 1}, want: []int{3, 2, 1}},
		{initStack: []int{3, 1, 2}, want: []int{3, 2, 1}},
		{initStack: []int{3, 2, 1}, want: []int{3, 2, 1}},
		{initStack: []int{3, 1, 4, 1, 5, 9, 2, 6, 5}, want: []int{9, 6, 5, 5, 4, 3, 2, 1, 1}},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("SortStack(%v)", tc.initStack)
		testFunc := func(t *testing.T) {
			stack := lib.ArrayStackFromSlice(tc.initStack)

			SortStack(stack)
			got := stack.ToSlice()

			if !slices.Equal(got, tc.want) {
				t.Fatalf("want %v, got %v", tc.want, got)
			} else {
				t.Logf("got %v", got)
			}
		}
		t.Run(testName, testFunc)
	}
}
