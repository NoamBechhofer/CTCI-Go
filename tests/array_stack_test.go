package lib

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
	"github.com/NoamBechhofer/CTCI-Go/testutils"
)

func TestArrayStack(t *testing.T) {
	stack := lib.ArrayStack[int]{}

	testutils.StackFunctionality(t, &stack)
}
