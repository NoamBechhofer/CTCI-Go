package lib

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
	"github.com/NoamBechhofer/CTCI-Go/testutils"
)

func TestArrayQueue(t *testing.T) {
	queue := lib.ArrayQueue[int]{}

	testutils.QueueFunctionality(t, &queue)
}
