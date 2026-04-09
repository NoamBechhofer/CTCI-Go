package minimaltree

import (
	"math"
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

type TestCase struct {
	uniqueIncreasing []int
}

func TestMinimalTree(t *testing.T) {
	testCases := []TestCase{{uniqueIncreasing: []int{}}}

	list := []int{}
	for i := 0; i < 1_000; i++ {
		list = append(list, i)
		testCases = append(testCases, TestCase{uniqueIncreasing: list})
	}

	for i, tc := range testCases {
		t.Run(lib.SignedToString(i), func(t *testing.T) {
			t.Log(tc.uniqueIncreasing)

			minimalTree := MinimalTree(tc.uniqueIncreasing)
			if !lib.IsSearchTree(minimalTree) {
				t.Fatalf("tree is not a valid search tree")
			}

			got := minimalTree.Height()
			var want int

			log2 := math.Log2(float64(len(tc.uniqueIncreasing)))
			if math.IsInf(log2, -1) {
				want = 0
			} else {
				want = 1 + int(math.Floor(log2))
			}

			if got != want {
				t.Fatalf("wanted %d, got %d", want, got)
			} else {
				t.Logf("got %d", got)
			}
		})
	}
}
