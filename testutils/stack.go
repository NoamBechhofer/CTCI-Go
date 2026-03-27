package testutils

import (
	"testing"

	"github.com/NoamBechhofer/CTCI-Go/lib"
)

func testStackIsEmpty[T any](t *testing.T, stack lib.Stack[T], expected bool) {
	t.Run("IsEmpty()", func(t *testing.T) {
		isEmpty := stack.IsEmpty()
		if expected != isEmpty {
			t.Fatalf("expected %t, got %t", expected, isEmpty)
		}
	})
}

func testStackPeek[T comparable](t *testing.T, stack lib.Stack[T], expectedVal T, expectedOk bool) {
	t.Run("Peek()", func(t *testing.T) {
		peeked, ok := stack.Peek()
		t.Run("correct emptiness report", func(t *testing.T) {
			if ok != expectedOk {
				t.Fatalf("expected %t, got %t", expectedOk, ok)
			}
		})
		t.Run("correct element peeked", func(t *testing.T) {
			if peeked != expectedVal {
				t.Fatalf("expected %v, got %v", expectedVal, peeked)
			}
		})
	})
}

func testStackPop[T comparable](t *testing.T, stack lib.Stack[T], expectedVal T, expectedOk bool) {
	t.Run("Pop()", func(t *testing.T) {
		popped, ok := stack.Pop()
		t.Run("correct emptiness report", func(t *testing.T) {
			if ok != expectedOk {
				t.Fatalf("expected %t, got %t", expectedOk, ok)
			}
		})
		t.Run("correct element popped", func(t *testing.T) {
			if popped != expectedVal {
				t.Fatalf("expected %v, got %v", expectedVal, popped)
			}
		})
	})
}

func testStackPush[T any](t *testing.T, stack lib.Stack[T], val T) {
	t.Run("Push()", func(t *testing.T) {
		stack.Push(val)
	})
}

func testStackSize[T any](t *testing.T, stack lib.Stack[T], expected int) {
	t.Run("Size()", func(t *testing.T) {
		got := stack.Size()
		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func testStackBasicOperationsAndEmpty(t *testing.T, stack lib.Stack[int]) {
	t.Run("basic stack operations test", func(t *testing.T) {
		testEmpty := func(t *testing.T) {
			testStackIsEmpty(t, stack, true)
			testStackPeek(t, stack, 0, false)
			testStackPop(t, stack, 0, false)
			testStackSize(t, stack, 0)
		}

		testPeek := func(expected int) func(t *testing.T) {
			return func(t *testing.T) {
				testStackIsEmpty(t, stack, false)
				testStackPeek(t, stack, expected, true)
			}
		}

		testPop := func(expected int) func(t *testing.T) {
			return func(t *testing.T) {
				testStackIsEmpty(t, stack, false)
				testStackPop(t, stack, expected, true)
			}
		}

		t.Run("empty", testEmpty)
		stack.Push(1)
		t.Run("has first element", testPeek(1))
		testStackSize(t, stack, 1)
		stack.Push(2)
		t.Run("has second element", testPeek(2))
		testStackSize(t, stack, 2)
		t.Run("pops second element", testPop(2))
		testStackSize(t, stack, 1)
		t.Run("pops first element", testPop(1))
		t.Run("empty again", testEmpty)
	})
}

func testStackPush10kAndEmpty(t *testing.T, stack lib.Stack[int]) {
	t.Run("10k pushes and 10k pops", func(t *testing.T) {
		for i := range 10_000 {
			stack.Push(i)
		}
		for i := 10_000 - 1; i >= 0; i-- {
			popped, ok := stack.Pop()
			if !ok {
				t.Fatalf("expected to pop %d, but stack was empty", i)
			}
			if popped != i {
				t.Fatalf("expected to pop %d, got %d", i, popped)
			}
		}
		testStackPop(t, stack, 0, false)
	})
}

func StackFunctionality(t *testing.T, stack lib.Stack[int]) {
	testStackBasicOperationsAndEmpty(t, stack)
	testStackPush10kAndEmpty(t, stack)
}
