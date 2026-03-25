package lib

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
	testutils "github.com/NoamBechhofer/CTCI-Go/tests/utils"
)

func TestArrayStack(t *testing.T) {
	stack := lib.ArrayStack[int]{}

	testutils.StackFunctionality(t, &stack)
}
