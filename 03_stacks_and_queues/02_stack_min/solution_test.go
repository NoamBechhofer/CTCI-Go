package stackmin

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/testutils"
)

func runIsEmpty(t *testing.T, stack *StackMin[int32], want bool) {
	t.Run("IsEmpty", func(t *testing.T) {
		got := stack.IsEmpty()
		if want != got {
			t.Fatalf("want %t, got %t", want, got)
		}
	})
}

func runPeek(t *testing.T, stack *StackMin[int32], wantVal int32, wantOK bool) {
	t.Run("Peek()", func(t *testing.T) {
		gotVal, gotOK := stack.Peek()
		if wantOK != gotOK || wantVal != gotVal {
			t.Fatalf("want (%d, %t), got (%d, %t)", wantVal, wantOK, gotVal, gotOK)
		}
	})
}

func runMin(t *testing.T, stack *StackMin[int32], wantVal int32, wantOK bool) {
	t.Run("Min()", func(t *testing.T) {
		gotVal, gotOK := stack.Min()
		if wantOK != gotOK || wantVal != gotVal {
			t.Fatalf("want (%d, %t), got (%d, %t)", wantVal, wantOK, gotVal, gotOK)
		}
	})
}

func runPop(t *testing.T, stack *StackMin[int32], wantVal int32, wantOK bool) {
	t.Run("Pop()", func(t *testing.T) {
		gotVal, gotOK := stack.Pop()
		if wantOK != gotOK || wantVal != gotVal {
			t.Fatalf("want (%d, %t), got (%d, %t)", wantVal, wantOK, gotVal, gotOK)
		}
	})
}

func TestStackMin(t *testing.T) {
	stackA := StackMin[int32]{}

	runIsEmpty(t, &stackA, true)
	t.Run("first Push()", func(t *testing.T) {
		stackA.Push(1)

		runPeek(t, &stackA, int32(1), true)
		runMin(t, &stackA, int32(1), true)
		runIsEmpty(t, &stackA, false)
	})

	t.Run("second Push()", func(t *testing.T) {
		stackA.Push(2)

		runPeek(t, &stackA, int32(2), true)
		runMin(t, &stackA, int32(1), true)
		runIsEmpty(t, &stackA, false)
	})

	t.Run("third Push()", func(t *testing.T) {
		stackA.Push(0)

		runPeek(t, &stackA, int32(0), true)
		runMin(t, &stackA, int32(0), true)
		runIsEmpty(t, &stackA, false)
	})

	t.Run("first Pop()", func(t *testing.T) {
		runPop(t, &stackA, int32(0), true)
		runPeek(t, &stackA, int32(2), true)
		runMin(t, &stackA, int32(1), true)
		runIsEmpty(t, &stackA, false)
	})

	t.Run("second Pop()", func(t *testing.T) {
		runPop(t, &stackA, int32(2), true)
		runPeek(t, &stackA, int32(1), true)
		runMin(t, &stackA, int32(1), true)
		runIsEmpty(t, &stackA, false)
	})

	t.Run("third Pop()", func(t *testing.T) {
		runPop(t, &stackA, int32(1), true)
		runPeek(t, &stackA, int32(0), false)
		runMin(t, &stackA, int32(0), false)
		runIsEmpty(t, &stackA, true)
	})

	stackB := StackMin[int]{}
	testutils.StackFunctionality(t, &stackB)
}
