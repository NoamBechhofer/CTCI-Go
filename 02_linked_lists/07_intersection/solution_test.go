package intersection

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	aVals           []int32
	bVals           []int32
	intersectionLen int
}

func intToString(num int32) string {
	return strconv.FormatInt(int64(num), 10)
}

func (tc *TestCase) toString() string {
	aUniqueVals := max(len(tc.aVals)-int(tc.intersectionLen), 0)
	bUniqueVals := max(len(tc.bVals)-int(tc.intersectionLen), 0)

	var aUniqueBuilder strings.Builder
	var bUniqueBuilder strings.Builder

	for i := range aUniqueVals {
		aUniqueBuilder.WriteString(intToString(tc.aVals[i]))
		if i != aUniqueVals-1 {
			aUniqueBuilder.WriteString(" → ")
		}
	}
	if tc.intersectionLen > 0 {
		aUniqueBuilder.WriteString(" ↘ ")
	}

	for i := range bUniqueVals {
		bUniqueBuilder.WriteString(intToString(tc.bVals[i]))
		if i != bUniqueVals-1 {
			bUniqueBuilder.WriteString(" → ")
		}
	}
	if tc.intersectionLen > 0 {
		bUniqueBuilder.WriteString(" ↗ ")
	}

	if tc.intersectionLen == 0 {
		var retBuilder strings.Builder
		retBuilder.WriteString(aUniqueBuilder.String())
		retBuilder.WriteString("\n")
		retBuilder.WriteString(bUniqueBuilder.String())
		return retBuilder.String()
	}

	var sharedBuilder strings.Builder
	for i := aUniqueVals; i < len(tc.aVals); i++ {
		sharedBuilder.WriteString(intToString(tc.aVals[i]))
		if i != len(tc.aVals)-1 {
			sharedBuilder.WriteString(" → ")
		}
	}

	aStr := aUniqueBuilder.String()
	bStr := bUniqueBuilder.String()

	uniqueLen := max(len([]rune(aStr)), len([]rune(bStr)))

	aUnique := strings.Repeat(" ", uniqueLen-len([]rune(aStr))) + aStr
	bUnique := strings.Repeat(" ", uniqueLen-len([]rune(bStr))) + bStr
	shared := strings.Repeat(" ", uniqueLen) + sharedBuilder.String()

	var retBuilder strings.Builder
	retBuilder.WriteString(aUnique)
	retBuilder.WriteString("\n")
	retBuilder.WriteString(shared)
	retBuilder.WriteString("\n")
	retBuilder.WriteString(bUnique)
	return retBuilder.String()
}

func (tc *TestCase) buildLists() (lib.SinglyLinkedList[int32], lib.SinglyLinkedList[int32], *lib.SinglyLinkedListNode[int32]) {
	var sharedHead *lib.SinglyLinkedListNode[int32] = nil
	for i := range tc.intersectionLen {
		sharedHead = &lib.SinglyLinkedListNode[int32]{Val: tc.aVals[len(tc.aVals)-1-i], Next: sharedHead}
	}
	aHead := sharedHead
	for i := len(tc.aVals) - 1 - tc.intersectionLen; i >= 0; i-- {
		aHead = &lib.SinglyLinkedListNode[int32]{Val: tc.aVals[i], Next: aHead}
	}

	bHead := sharedHead
	for i := len(tc.bVals) - 1 - tc.intersectionLen; i >= 0; i-- {
		bHead = &lib.SinglyLinkedListNode[int32]{Val: tc.bVals[i], Next: bHead}
	}

	return lib.SinglyLinkedList[int32]{Head: aHead}, lib.SinglyLinkedList[int32]{Head: bHead}, sharedHead
}

func TestIntersection(t *testing.T) {
	testCases := []TestCase{
		{aVals: []int32{}, bVals: []int32{}, intersectionLen: 0},
		{aVals: []int32{}, bVals: []int32{1}, intersectionLen: 0},
		{aVals: []int32{1}, bVals: []int32{1}, intersectionLen: 0},
		{aVals: []int32{1}, bVals: []int32{1}, intersectionLen: 1},
		{aVals: []int32{1, 2}, bVals: []int32{2}, intersectionLen: 1},
		{aVals: []int32{1, 2}, bVals: []int32{2}, intersectionLen: 0},
		{aVals: []int32{1, 2}, bVals: []int32{1, 2}, intersectionLen: 0},
		{aVals: []int32{1, 2}, bVals: []int32{1, 2}, intersectionLen: 1},
		{aVals: []int32{1, 2}, bVals: []int32{1, 2}, intersectionLen: 2},
		{aVals: []int32{1, 2, 3, 4, 5}, bVals: []int32{3, 2, 1, 4, 5}, intersectionLen: 2},
		{aVals: []int32{1, 2, 3, 4, 5}, bVals: []int32{3, 2, 1, 4, 5}, intersectionLen: 0},
		{aVals: []int32{1, 2, 3}, bVals: []int32{9, 2, 3}, intersectionLen: 0},
		{aVals: []int32{1, 2, 3, 4, 5, 6}, bVals: []int32{9, 6}, intersectionLen: 1},
		{aVals: []int32{1, 2, 3, 4, 5, 6}, bVals: []int32{9, 8, 5, 6}, intersectionLen: 2},
		{aVals: []int32{1, 2, 3, 4}, bVals: []int32{3, 4}, intersectionLen: 2},
		{aVals: []int32{1, 2, 3}, bVals: []int32{1, 2, 3}, intersectionLen: 3},
		{aVals: []int32{1, 2, 1, 2, 3}, bVals: []int32{9, 1, 2, 3}, intersectionLen: 3},
		{aVals: []int32{1, 1, 1, 2}, bVals: []int32{3, 1, 2}, intersectionLen: 2},
		{aVals: []int32{1, 2, 7}, bVals: []int32{3, 4, 7}, intersectionLen: 0},
		{aVals: []int32{1, 2, 3}, bVals: []int32{1, 2, 3}, intersectionLen: 3},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("Intersection(%v, %v, %d)", tc.aVals, tc.bVals, tc.intersectionLen)
		testFunc := func(t *testing.T) {
			t.Logf("\n%s\n", tc.toString())
			a, b, expected := tc.buildLists()

			got := Intersection(a, b)

			var gotVal string
			if got != nil {
				gotVal = intToString(got.Val)
			} else {
				gotVal = "nil"
			}

			if got != expected {
				var expectedVal string
				if expected != nil {
					expectedVal = intToString(expected.Val)
				} else {
					expectedVal = "nil"
				}

				t.Fatalf("expected %p (%s), got %p (%s)", expected, expectedVal, got, gotVal)
			} else {
				t.Logf("got %p (%s)", got, gotVal)
			}
		}
		t.Run(testName, testFunc)
	}
}
