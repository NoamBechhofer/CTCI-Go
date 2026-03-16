package sumlists

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	a uint32
	b uint32
}

// given nonnegative integer n, return slices contianing the digits: the first
// return contains the digits in forward order, the second contains them in
// reverese order. e.g. createSlices(1234) = ([1, 2, 3, 4], [4, 3, 2, 1])
func createSlices(n uint32) ([]uint32, []uint32) {
	float := float64(n)
	log10 := int(math.Log10(float))

	forward := []uint32{}
	reverse := []uint32{}

	for placeIdx := log10; placeIdx >= 0; placeIdx-- {
		place := uint32(math.Pow10(placeIdx))
		digit := n / place
		forward = append(forward, digit)
		reverse = append([]uint32{digit}, reverse...)
		n %= place
	}

	return forward, reverse
}

func TestSumLists(t *testing.T) {
	testCases := []TestCase{
		{a: 0, b: 0},
		{a: 1, b: 11},
		{a: 1, b: 111},
		{a: 0, b: 123},
		{a: 0, b: 321},
		{a: 617, b: 295},
		{a: 716, b: 592},
		{a: 123, b: 123},
		{a: 321, b: 321},
		{a: 183, b: 352},
		{a: 381, b: 253},
		{a: 183, b: 952},
		{a: 381, b: 259},
		{a: 189, b: 952},
		{a: 981, b: 259},
		{a: 189, b: 52},
		{a: 981, b: 25},
	}

	for _, tc := range testCases {

		aForward, aReverse := createSlices(tc.a)
		bForward, bReverse := createSlices(tc.b)
		expected := tc.a + tc.b
		expectedForward, expectedReverse := createSlices(expected)

		t.Run(fmt.Sprintf("SumListsReverse(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			aList := lib.SinglyLinkedListFromSlice(aReverse)
			bList := lib.SinglyLinkedListFromSlice(bReverse)
			got := SumListsReverse(aList, bList)
			gotSlice := got.ToSlice()
			if !slices.Equal(gotSlice, expectedReverse) {
				t.Fatalf("expected %d, got %v", expected, gotSlice)
			} else {
				t.Logf("got %v", gotSlice)
			}
		})

		t.Run(fmt.Sprintf("SumListsForward(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			aList := lib.SinglyLinkedListFromSlice(aForward)
			bList := lib.SinglyLinkedListFromSlice(bForward)
			got := SumListsForward(aList, bList)
			gotSlice := got.ToSlice()
			if !slices.Equal(gotSlice, expectedForward) {
				t.Fatalf("expected %d, got %v", expected, gotSlice)
			} else {
				t.Logf("got %v", gotSlice)
			}
		})
	}
}
