package misc

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func TestFactorialRecursive(t *testing.T) {
	for i := range uint(1000) {
		t.Run(lib.UnsignedToString(i), func(t *testing.T) {
			if FactorialRecursive(i) != FactorialIterative(i) {
				t.Fail()
			}
		})
	}
}
