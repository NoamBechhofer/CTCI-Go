package stackofplates

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/testutils"
)

func TestStackOfPlates(t *testing.T) {
	stackA := NewStackOfPlates[int]()

	testutils.StackFunctionality(t, stackA)

	stackB := NewStackOfPlates[int]()
	for i := range 10 * MAX_STACK_SIZE {
		stackB.Push(i)
	}

	pop, popOk := stackB.PopAt(0)
	if !popOk {
		t.Fatalf("expected PopAt to succeed")
	}
	want := MAX_STACK_SIZE - 1
	if pop != want {
		t.Fatalf("expected PopAt to return %d, got %d", want, pop)
	}

	pop, popOk = stackB.PopAt(5)
	if !popOk {
		t.Fatalf("expected PopAt to succeed")
	}
	want = (5 * MAX_STACK_SIZE) + MAX_STACK_SIZE
	if pop != want {
		t.Fatalf("expected PopAt to return %d, got %d", want, pop)
	}
}
