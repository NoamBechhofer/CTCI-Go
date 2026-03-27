package queueviastacks

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/testutils"
)

func TestQueueViaStacks(t *testing.T) {
	var queue QueueViaStacks[int]
	testutils.QueueFunctionality(t, &queue)
}
